package app

import (
	"context"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/cmd/config"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/handlers"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ServerApp struct {
	cfg    config.ServerConfig
	logger *loggers.Logger
	router *mux.Router
}

func NewServerApp() *ServerApp {
	cfg := config.ServerConfigInit()
	router := mux.NewRouter()
	logger := loggers.NewLogger()

	return &ServerApp{
		router: router,
		cfg:    *cfg,
		logger: logger,
	}
}

func (a *ServerApp) Run() {
	client, err := postgres.NewClient(context.Background(), 5, &a.cfg, a.logger)
	if err != nil {
		log.Err(err)
		return
	}
	store, err := repositories.NewPGStore(client, &a.cfg, a.logger)
	if err != nil {
		log.Err(err)
		return
	}

	h := handlers.NewHandler(&a.cfg, a.logger, store)

	h.Register(a.router)

	server := &http.Server{}
	server.Handler = a.router
	server.Addr = a.cfg.Addr

	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return
	}

}
