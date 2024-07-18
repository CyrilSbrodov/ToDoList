package service

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
)

// Service - интерфейс транспорта
type Service interface {
	CreateUser(ctx context.Context, u *models.User) (string, error)
	GenerateToken(ctx context.Context, u *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
	NewTask(ctx context.Context, list *models.TodoList) error
	GetAll(ctx context.Context, u *models.User) error
	CreateGroup(ctx context.Context, list *models.TodoList) error
	AddUserToGroup(ctx context.Context, list *models.TodoList) error
	DeleteGroup(ctx context.Context, list *models.TodoList) error
	RemoveUserFromGroup(ctx context.Context, list *models.TodoList) error
}
