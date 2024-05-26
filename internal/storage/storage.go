package storage

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
)

type Storage interface {
	NewUser(ctx context.Context, u *models.User) (string, error)
	Auth(ctx context.Context, u *models.User) (string, error)
	NewTask(ctx context.Context, list *models.TodoList) error
	GetAll(ctx context.Context, u *models.User) error
}
