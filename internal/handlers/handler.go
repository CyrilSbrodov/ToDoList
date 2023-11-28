package handlers

import (
	"github.com/CyrilSbrodov/ToDoList/cmd/config"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Handlers interface {
	Register(router *chi.Mux)
}

type Handler struct {
	cfg     *config.ServerConfig
	logger  *loggers.Logger
	storage storage.Storage
}

func NewHandler(cfg *config.ServerConfig, logger *loggers.Logger, storage storage.Storage) *Handler {
	return &Handler{
		cfg:     cfg,
		logger:  logger,
		storage: storage,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	//r.Group(func(r chi.Router) {
	//	r.Post("/api/register", h.Registration())
	//	r.Post("/api/login", h.Login())
	//})
	r.Group(func(r chi.Router) {
		r.Get("/api/", h.GetAll())
		r.Post("/api/new", h.NewList())
		//r.Post("/api/group", h.AddGroup())
	})
}
