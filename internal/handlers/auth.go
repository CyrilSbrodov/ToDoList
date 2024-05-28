package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		id, err := h.service.CreateUser(r.Context(), &user)
		if err != nil {
			if errors.Is(err, models.ErrorUserConflict) {
				http.Error(w, "User name or email already exists", http.StatusConflict)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User: %s has been created. ID : %s", user.Name, id)
		return
	}
}

func (h *Handler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		token, err := h.service.GenerateToken(r.Context(), &u)
		if err != nil {
			if errors.Is(err, models.ErrorUserNotFound) {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, token)
		return
	}
}
