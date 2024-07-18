package handlers

import (
	"encoding/json"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"net/http"
)

func (h *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		user.Id = r.Context().Value(ctxKeyUser).(string)
		//err := h.service.GetAll(r.Context(), &user)
		//// TODO err no rows and internal
		//if err != nil {
		//	w.Header().Set("Content-Type", "application/json")
		//	http.Error(w, "empty task", http.StatusBadRequest)
		//	return
		//}
		data, err := json.Marshal(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "failed to encode data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}
