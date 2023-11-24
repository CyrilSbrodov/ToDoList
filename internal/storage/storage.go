package storage

import "github.com/CyrilSbrodov/ToDoList/internal/storage/models"

type Storage interface {
	GetAll() models.User
	NewList(user models.User) error
}
