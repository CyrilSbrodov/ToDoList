package storage

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
)

// Storage - интерфейс БД
type Storage interface {
	NewUser(ctx context.Context, u *models.User) (string, error)
	Auth(ctx context.Context, u *models.User) (string, error)
	NewTask(ctx context.Context, list *models.TodoList) error
	GetAll(ctx context.Context, u *models.User) error
	CreateGroup(ctx context.Context, list *models.TodoList) error
	AddUserToGroup(ctx context.Context, list *models.TodoList) error
	DeleteGroup(ctx context.Context, list *models.TodoList) error
	RemoveUserFromGroup(ctx context.Context, list *models.TodoList) error
}
