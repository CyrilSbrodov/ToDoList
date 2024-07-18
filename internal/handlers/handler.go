package handlers

import (
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/service"
	"github.com/gorilla/mux"
)

// Handlers - интерфейс хэндлера (возможно стоит вынести в другое место)
type Handlers interface {
	Register(router *mux.Router)
}

// Handler - структура хэндлера
type Handler struct {
	cfg     *config.Config
	logger  *loggers.Logger
	service service.Service
}

// NewHandler - конструктор хэндлера
func NewHandler(cfg *config.Config, logger *loggers.Logger, service service.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		logger:  logger,
		service: service,
	}
}

// Register - метод регистрации эндпоинтов
func (h *Handler) Register(r *mux.Router) {
	r.HandleFunc("/api/register", h.signUp()).Methods("POST")
	r.HandleFunc("/api/login", h.signIn()).Methods("POST")
	secure := r.PathPrefix("/auth").Subrouter()
	secure.Use(h.userIdentity)
	secure.HandleFunc("/api/task", h.getAll()).Methods("GET")
	secure.HandleFunc("/api/task", h.newList()).Methods("POST")
	secure.HandleFunc("/api/groups", h.createGroup()).Methods("POST")
	secure.HandleFunc("/api/groups/{groupId}/users", h.addUserToGroup()).Methods("POST")
	secure.HandleFunc("/api/groups/{groupId}", h.deleteGroup()).Methods("DELETE")
	secure.HandleFunc("/api/groups/{groupId}/users/{userId}", h.removeUserFromGroup()).Methods("DELETE")
}
