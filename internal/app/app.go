package app

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/handlers"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/gorilla/mux"
	"net/http"
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
		a.logger.LogErr(err, "failed to start pg client")
		return
	}
	store, err := repositories.NewPGStore(client, &a.cfg, a.logger)
	if err != nil {
		a.logger.LogErr(err, "failed to start pg store")
		return
	}

	h := handlers.NewHandler(&a.cfg, a.logger, store)

	h.Register(a.router)

	server := &http.Server{
		Addr:         a.cfg.Listener.Addr,
		Handler:      a.router,
		ReadTimeout:  a.cfg.Listener.Timeout,
		WriteTimeout: a.cfg.Listener.Timeout,
		IdleTimeout:  a.cfg.Listener.IdleTimeout,
	}

	if err = server.ListenAndServe(); err != nil {
		a.logger.LogErr(err, "failed to start server")
		return
	}
}
