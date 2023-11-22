package app

import (
	"github.com/CyrilSbrodov/ToDoList/cmd/config"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/go-chi/chi/v5"
)

type ServerApp struct {
	cfg    config.ServerConfig
	logger *loggers.Logger
	router *chi.Mux
}

func NewServerApp() *ServerApp {
	cfg := config.ServerConfigInit()
	router := chi.NewRouter()
	logger := loggers.NewLogger()

	return &ServerApp{
		router: router,
		cfg:    *cfg,
		logger: logger,
	}
}

func (a *ServerApp) Run() {

}
