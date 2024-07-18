package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) NewList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(ctxKeyUser).(string)
		fmt.Println("id: ", id)
		if id == "" {
			fmt.Println("error, nil id")
			return
		}
		var list models.TodoList
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		fmt.Println(list)
		// TODO проверка полей юзера
		if err := h.service.NewGroup(r.Context(), &list); err != nil {
			h.logger.Error("func newlist", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
