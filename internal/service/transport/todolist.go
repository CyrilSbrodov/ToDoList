package transport

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
)

func (t *Transport) NewTask(ctx context.Context, list *models.TodoList) error {
	return nil
}

func (t *Transport) GetAll(ctx context.Context, u *models.User) error {
	return nil
}

func (t *Transport) CreateGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.CreateGroup(ctx, list)
}
func (t *Transport) AddUserToGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.AddUserToGroup(ctx, list)
}
func (t *Transport) DeleteGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.DeleteGroup(ctx, list)
}
func (t *Transport) RemoveUserFromGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.RemoveUserFromGroup(ctx, list)
}
