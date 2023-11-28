package handlers

import (
	"encoding/json"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/models"
	"io"
	"net/http"
)

func (h *Handler) NewList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var user models.User
		if err = json.Unmarshal(content, &user); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO проверка полей юзера
		if err = h.storage.NewList(user); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
