package storage

import "github.com/CyrilSbrodov/ToDoList/internal/storage/models"

type Storage interface {
	NewUser(user *models.User) error
	Auth(user *models.User) error
	NewList(user *models.User) error
	GetAll(user *models.User) (models.User, error)
}
