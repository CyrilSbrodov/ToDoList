package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) Registration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ToDo check auth
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		id, err := h.storage.NewUser(r.Context(), &user)
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
