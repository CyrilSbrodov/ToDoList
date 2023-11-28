package repositories

import (
	"context"
	"github.com/CyrilSbrodov/ToDoList/cmd/config"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/storage/models"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/jackc/pgx/v5"
	"time"
)

type PGStore struct {
	client postgres.Client
	cfg    *config.ServerConfig
	logger *loggers.Logger
}

// createTable - функция создания новых таблиц в БД.
func createTable(ctx context.Context, client postgres.Client, logger *loggers.Logger) error {
	tx, err := client.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.LogErr(err, "failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	//создание таблиц
	q := `CREATE TABLE if not exists users (
    		id BIGINT PRIMARY KEY generated always as identity,
    		login VARCHAR(200) NOT NULL unique,
    		hashed_password VARCHAR(200) NOT NULL
		);
		CREATE UNIQUE INDEX if not exists users_login_uindex on users (login);
		CREATE TABLE if not exists user_group (
    		user_id BIGINT,
    		id BIGINT PRIMARY KEY generated always as identity,
    		FOREIGN KEY (user_id) REFERENCES users(id),
    		text bytea                         
		);
		CREATE TABLE if not exists binary_table (
    		user_id BIGINT,
    		id BIGINT PRIMARY KEY generated always as identity,
    		FOREIGN KEY (user_id) REFERENCES users(id),
    		binary_data bytea                          
		);
		CREATE TABLE if not exists passwords (
    		user_id BIGINT,
    		id BIGINT PRIMARY KEY generated always as identity,
    		FOREIGN KEY (user_id) REFERENCES users(id),
    		login bytea,
		    password bytea
		);
		CREATE TABLE if not exists cards (
    		user_id BIGINT,
    		id BIGINT PRIMARY KEY generated always as identity,
    		card_number bytea,
    		FOREIGN KEY (user_id) REFERENCES users(id),
    		card_holder bytea,
    		cvc bytea                            
		);
		CREATE UNIQUE INDEX if not exists cards_card_number_uindex on cards (card_number);`

	_, err = tx.Exec(ctx, q)
	if err != nil {
		logger.LogErr(err, "failed to create table")
		return err
	}
	return tx.Commit(ctx)
}

func NewPGStore(client postgres.Client, cfg *config.ServerConfig, logger *loggers.Logger) (*PGStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := createTable(ctx, client, logger); err != nil {
		logger.LogErr(err, "failed to create table")
		return nil, err
	}
	return &PGStore{
		client: client,
		cfg:    cfg,
		logger: logger,
	}, nil
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
