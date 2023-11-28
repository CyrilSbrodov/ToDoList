package handlers

import (
	"encoding/json"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/models"
	"net/http"
)

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//ToDo check auth
		user := models.User{Id: "1"}
		user, err := h.storage.GetAll(&user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		userJson, err := json.Marshal(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	}
}
