package transport

import (
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/storage"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
)

// Transport - структура слоя транспорт
type Transport struct {
	storage.Storage
	logger *loggers.Logger
}

// NewTransport - конструктор транспорта
func NewTransport(repo repositories.PGStore, logger *loggers.Logger) *Transport {
	return &Transport{
		&repo,
		logger,
	}
}
