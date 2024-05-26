package handlers

import (
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/service"
	"github.com/gorilla/mux"
)

type Handlers interface {
	Register(router *mux.Router)
}

type Handler struct {
	cfg     *config.Config
	logger  *loggers.Logger
	service service.Service
}

func NewHandler(cfg *config.Config, logger *loggers.Logger, service service.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		logger:  logger,
		service: service,
	}
}

func (h *Handler) Register(r *mux.Router) {
	//r.Group(func(r chi.Router) {
	//	r.Post("/api/register", h.Registration())
	//	r.Post("/api/login", h.Login())
	//})
	r.HandleFunc("/api/register", h.SignUp()).Methods("POST")
	r.HandleFunc("/api/login", h.SignIn()).Methods("POST")
	secure := r.PathPrefix("/auth").Subrouter()
	//secure.Use(JwtVerify)
	secure.HandleFunc("/api/task", h.GetAll()).Methods("GET")
	secure.HandleFunc("/api/task", h.NewList()).Methods("POST")
	//r.Group(func(r chi.Router) {
	//	r.Get("/api/", h.GetAll())
	//	r.Post("/api/new", h.NewList())
	//	r.Post("/api/group", h.AddGroup())
	//})
}
