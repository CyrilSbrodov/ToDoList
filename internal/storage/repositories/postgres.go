package repositories

import "github.com/CyrilSbrodov/ToDoList/internal/storage/models"

type PGStore struct {
}

func NewPGStore() *PGStore {
	return &PGStore{}
}

func (p *PGStore) NewUser(u *models.User) error {
	return nil
}

func (p *PGStore) Auth(u *models.User) error {
	return nil
}

func (p *PGStore) NewList(u *models.User) error {
	return nil
}

func (p *PGStore) GetAll(u *models.User) (models.User, error) {
	return *u, nil
}
