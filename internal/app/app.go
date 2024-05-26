package app

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/handlers"
	"github.com/CyrilSbrodov/ToDoList/internal/service/transport"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
)

type ServerApp struct {
	cfg    config.Config
	logger *loggers.Logger
	router *mux.Router
}

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

func (a *ServerApp) Run() {
	client, err := postgres.NewClient(context.Background(), 5, &a.cfg, a.logger)
	if err != nil {
		a.logger.Error("failed to start pg client", err)
		return
	}

	store, err := repositories.NewPGStore(client, &a.cfg, a.logger)
	if err != nil {
		a.logger.Error("failed to start pg store", err)
		return
	}
	t := transport.NewTransport(*store)
	h := handlers.NewHandler(&a.cfg, a.logger, t)

	h.Register(a.router)

	srv := &http.Server{
		Addr:         a.cfg.Listener.Addr,
		Handler:      a.router,
		ReadTimeout:  a.cfg.Listener.Timeout,
		WriteTimeout: a.cfg.Listener.Timeout,
		IdleTimeout:  a.cfg.Listener.IdleTimeout,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			a.logger.Error("server", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		a.logger.Info("server", "failed to shutting down gracefully")
		return
	}
	a.logger.Info("server", "shutting down")
	os.Exit(0)
}
