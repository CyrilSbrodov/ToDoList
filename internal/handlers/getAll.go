package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//ToDo check auth

		user := h.storage.GetAll()
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
