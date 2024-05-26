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
