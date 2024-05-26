package transport

import (
	"github.com/CyrilSbrodov/ToDoList/internal/storage"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/repositories"
)

type Transport struct {
	storage.Storage
}

func NewTransport(repo repositories.PGStore) *Transport {
	return &Transport{
		&repo,
	}
}
