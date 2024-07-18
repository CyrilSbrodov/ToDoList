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

func (t *Transport) NewGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.NewGroup(ctx, list)
}
func (t *Transport) AddInGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.AddInGroup(ctx, list)
}
func (t *Transport) DeleteGroup(ctx context.Context, list *models.TodoList) error {
	return t.Storage.DeleteGroup(ctx, list)
}
