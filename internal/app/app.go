package app

import (
	"context"
	"errors"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/handlers"
	"github.com/CyrilSbrodov/ToDoList/internal/service/transport"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ServerApp - структура сервера
type ServerApp struct {
	cfg    config.Config
	logger *loggers.Logger
	router *mux.Router
}

// NewServerApp - конструктор сервера
func NewServerApp() *ServerApp {
	cfg := config.NewConfig()
	router := mux.NewRouter()
	logger := loggers.SetupLogger(cfg.Env)

	return &ServerApp{
		router: router,
		cfg:    *cfg,
		logger: logger,
	}
}

// Run - метод запуска сервера
func (a *ServerApp) Run() {
	// создание нового клиента постгрес
	client, err := postgres.NewClient(context.Background(), 5, &a.cfg, a.logger)
	if err != nil {
		a.logger.Error("failed to start pg client", err, "server")
		return
	}
	// создание БД
	store, err := repositories.NewPGStore(client, &a.cfg, a.logger)
	if err != nil {
		a.logger.Error("failed to start pg store", err, "server")
		return
	}
	// инициализация транспорта
	t := transport.NewTransport(*store, a.logger)
	// инициализация хэндлеров
	h := handlers.NewHandler(&a.cfg, a.logger, t)
	// регистрация хэндлеров
	h.Register(a.router)

	srv := &http.Server{
		Addr:         a.cfg.Listener.Addr,
		Handler:      a.router,
		ReadTimeout:  a.cfg.Listener.Timeout,
		WriteTimeout: a.cfg.Listener.Timeout,
		IdleTimeout:  a.cfg.Listener.IdleTimeout,
	}
	// зпуск сервера в отдельной горутине
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			a.logger.Error("server not started", err, "server")
		}
	}()
	// создание канала для прослушивания сигналов прерывания
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// если сигнал получен, то начинается выключение сервера
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// отключение сервера
	if err = srv.Shutdown(ctx); err != nil {
		a.logger.Error("failed to shutting down gracefully", err, "server")
		return
	}
	a.logger.Info("shutting down", slog.String("server", a.cfg.Listener.Addr))
	os.Exit(0)
}
