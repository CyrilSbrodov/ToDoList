package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) NewGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.TodoList
		list.UserID = r.Context().Value(ctxKeyUser).(string)

		if list.UserID == "" {
			h.logger.Error("ID", errors.New("id is nil"))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			h.logger.Error("failed to decode json", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		// TODO проверка полей юзера
		if err := h.service.NewGroup(r.Context(), &list); err != nil {
			h.logger.Error("func NewGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (h *Handler) AddInGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.TodoList
		list.UserID = r.Context().Value(ctxKeyUser).(string)

		if list.UserID == "" {
			h.logger.Error("ID", errors.New("id is nil"))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if err := h.service.AddInGroup(r.Context(), &list); err != nil {
			h.logger.Error("func AddInGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (h *Handler) DeleteGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.TodoList
		list.UserID = r.Context().Value(ctxKeyUser).(string)

		if list.UserID == "" {
			h.logger.Error("ID", errors.New("id is nil"))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			//TODO log error
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		fmt.Println(list)
		// TODO проверка полей юзера
		if err := h.service.DeleteGroup(r.Context(), &list); err != nil {
			h.logger.Error("func DeleteGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
