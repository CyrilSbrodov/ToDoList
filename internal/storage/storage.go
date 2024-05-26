package storage

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
)

type Storage interface {
	NewUser(ctx context.Context, user *models.User) (string, error)
	Auth(ctx context.Context, user *models.User) error
	NewList(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context, user *models.User) (models.User, error)
}
