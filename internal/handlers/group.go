package handlers

import (
	"encoding/json"
	"errors"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

// createGroup - метод получения данных от пользователя для создания новой группы
func (h *Handler) createGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.TodoList
		//проверка id юзера из контекста
		list.UserID = r.Context().Value(ctxKeyUser).(string)
		//если айди не получен, то возвращаем ошибку
		if list.UserID == "" {
			h.logger.Error("ID", errors.New("id is nil"))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		//получение данных из body json
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			h.logger.Error("failed to decode json", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		//отправка данных ниже, в слой транспорт
		if err := h.service.CreateGroup(r.Context(), &list); err != nil {
			h.logger.Error("func NewGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// addUserToGroup - метод получения данных от пользователя для добавления нового пользователя в группу
func (h *Handler) addUserToGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получение данных из url
		vars := mux.Vars(r)
		var list models.TodoList
		//проверка id юзера из контекста
		list.UserID = r.Context().Value(ctxKeyUser).(string)
		//если айди не получен, то возвращаем ошибку
		if list.UserID == "" {
			h.logger.Error("ID", errors.New("id is nil"))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		//получение groupId из url
		list.GroupID = vars["groupId"]
		//получение данных из body json
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			h.logger.Error("func h.addUserToGroup", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		//отправка данных на слой ниже, в транспорт
		if err := h.service.AddUserToGroup(r.Context(), &list); err != nil {
			h.logger.Error("func AddInGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// deleteGroup - метод получения данных от пользователя для удаления группы
func (h *Handler) deleteGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получение данных из url
		vars := mux.Vars(r)
		var list models.TodoList
		//проверка id юзера из контекста (не уверен, что тут нужно)
		//list.UserID = r.Context().Value(ctxKeyUser).(string)
		//if list.UserID == "" {
		//	h.logger.Error("ID", errors.New("id is nil"))
		//	http.Error(w, "internal server error", http.StatusInternalServerError)
		//	return
		//}
		//получение данных из body json
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			h.logger.Error("func h.deleteGroup", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		//получение groupId из url
		list.GroupID = vars["groupId"]

		//отправка данных на слой ниже, в транспорт
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

// removeUserFromGroup - метод получения данных от пользователя для удаления пользователя из группы
func (h *Handler) removeUserFromGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получение данных из url
		vars := mux.Vars(r)
		var list models.TodoList
		//list.UserID = r.Context().Value(ctxKeyUser).(string)
		//if list.UserID == "" {
		//	h.logger.Error("ID", errors.New("id is nil"))
		//	http.Error(w, "internal server error", http.StatusInternalServerError)
		//	return
		//}

		//получение groupId и userId из url
		list.GroupID = vars["groupId"]
		list.AnotherUserID = vars["userId"]
		//получение данных из body json
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			h.logger.Error("func h.removeUserFromGroup", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		//отправка данных на слой ниже, в транспорт
		if err := h.service.RemoveUserFromGroup(r.Context(), &list); err != nil {
			h.logger.Error("func DeleteGroup", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
