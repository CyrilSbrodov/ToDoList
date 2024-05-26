package handlers

import (
	"encoding/json"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) NewList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.TodoList
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// TODO проверка полей юзера
		if err := h.service.NewTask(r.Context(), &list); err != nil {
			h.logger.Error("func newlist", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
