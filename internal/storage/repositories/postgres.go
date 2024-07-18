package repositories

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/cmd/loggers"
	"github.com/CyrilSbrodov/ToDoList/internal/config"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"github.com/CyrilSbrodov/ToDoList/pkg/client/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

// PGStore -структура базы
type PGStore struct {
	client postgres.Client
	cfg    *config.Config
	logger *loggers.Logger
}

// createTable - функция создания новых таблиц в БД.
func createTable(ctx context.Context, client postgres.Client, logger *loggers.Logger) error {
	tx, err := client.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error("failed to begin transaction", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	//создание таблиц
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
            user_id SERIAL PRIMARY KEY,
            user_name VARCHAR(50) NOT NULL UNIQUE,
            password_hash VARCHAR(255) NOT NULL,
            email VARCHAR(100) NOT NULL UNIQUE
        )`,
		`CREATE TABLE IF NOT EXISTS groups (
            group_id SERIAL PRIMARY KEY,
            group_name VARCHAR(100) NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS user_groups (
            user_id INT NOT NULL,
            group_id INT NOT NULL,
            PRIMARY KEY (user_id, group_id),
            FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
            FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
        )`,
		`CREATE TABLE IF NOT EXISTS tasks (
            task_id SERIAL PRIMARY KEY,
            task_description TEXT NOT NULL,
            created_by INT NOT NULL,
            group_id INT,
            is_completed BOOLEAN NOT NULL DEFAULT FALSE,
            FOREIGN KEY (created_by) REFERENCES users (user_id) ON DELETE SET NULL,
            FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE SET NULL
        )`,
		`CREATE TABLE IF NOT EXISTS task_assignees (
            task_id INT NOT NULL,
            user_id INT NOT NULL,
            PRIMARY KEY (task_id, user_id),
            FOREIGN KEY (task_id) REFERENCES tasks (task_id) ON DELETE CASCADE,
            FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
        )`,
	}

	for _, table := range tables {
		_, err = tx.Exec(ctx, table)
		if err != nil {
			logger.Error("Unable to create table", err)
			return err
		}
	}
	return tx.Commit(ctx)
}

// NewPGStore - конструктор БД
func NewPGStore(client postgres.Client, cfg *config.Config, logger *loggers.Logger) (*PGStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := createTable(ctx, client, logger); err != nil {
		logger.Error("failed to create table", err)
		return nil, err
	}
	return &PGStore{
		client: client,
		cfg:    cfg,
		logger: logger,
	}, nil
}

// NewUser - метод создания нового пользователя
func (p *PGStore) NewUser(ctx context.Context, u *models.User) (string, error) {
	//хеширование пароля
	hashPassword := p.hashPassword(u.Password)

	q := `INSERT INTO users (user_name, password_hash, email) VALUES ($1, $2, $3) RETURNING user_id`
	if err := p.client.QueryRow(ctx, q, u.Name, hashPassword, u.Email).Scan(&u.Id); err != nil {
		var pgErr *pgconn.PgError
		// проверка на ошибку конфликта из БД, если конфликт имени или почты, то возвращает ошибку из моделс
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return u.Id, models.ErrorUserConflict
		}
		p.logger.Error("Failure to insert object into table", err)
		return u.Id, err
	}
	return u.Id, nil

}

// Auth - метод авторизации пользователя с проверкой логина и пароля
func (p *PGStore) Auth(ctx context.Context, u *models.User) (string, error) {
	hashPassword := p.hashPassword(u.Password)
	q := `SELECT user_id FROM users WHERE user_name=$1 AND password_hash=$2`
	if err := p.client.QueryRow(ctx, q, u.Name, hashPassword).Scan(&u.Id); err != nil {
		// проверка на ошибку из БД, что не найден такой пользователь
		if errors.Is(err, pgx.ErrNoRows) {
			return "", models.ErrorUserNotFound
		}
		p.logger.Error("Failure to select object from table", err)
		return "", err
	}
	return u.Id, nil
}

// NewTask - функция создания новой задачи
func (p *PGStore) NewTask(ctx context.Context, list *models.TodoList) error {
	q := `INSERT INTO tasks (task_description, created_by, group_id, is_completed) VALUES ($1, $2, $3, $4)`
	if _, err := p.client.Exec(ctx, q, list.Task, list.UserID, list.Group, list.Status); err != nil {
		p.logger.Error("Failure to select object from table", err)
		return err
	}
	return nil
}

// CreateGroup - метод создания новой группы, где уникально только id группы. Название может быть любым.
// Идет создание группы в таблице groups и сразу же добавление id пользователя, который создал эту группу к группе в таблицу user_groups
func (p *PGStore) CreateGroup(ctx context.Context, list *models.TodoList) error {
	q := `INSERT INTO groups (group_name) VALUES ($1) RETURNING group_id`
	var groupID string
	if err := p.client.QueryRow(ctx, q, list.Group).Scan(&groupID); err != nil {
		//TODO err
		return err
	}
	q = `INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2)`
	if _, err := p.client.Exec(ctx, q, list.UserID, groupID); err != nil {
		//TODO err
		return err
	}
	return nil
}

// AddUserToGroup - добавление пользователя к группе по id группы и id пользователя
func (p *PGStore) AddUserToGroup(ctx context.Context, list *models.TodoList) error {
	q := `INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2)`
	if _, err := p.client.Exec(ctx, q, list.AnotherUserID, list.GroupID); err != nil {
		//TODO err
		return err
	}
	return nil
}

// DeleteGroup - метод удаления группы из двух таблиц путем транзакции
func (p *PGStore) DeleteGroup(ctx context.Context, list *models.TodoList) error {
	tx, err := p.client.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		//TODO err
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	q := `DELETE FROM groups WHERE group_id=$1`
	if _, err = tx.Exec(ctx, q, list.GroupID); err != nil {
		//TODO err
		return err
	}
	q = `DELETE FROM user_groups WHERE group_id=$1`
	if _, err = tx.Exec(ctx, q, list.GroupID); err != nil {
		//TODO err
		return err
	}
	return tx.Commit(ctx)
}

// RemoveUserFromGroup - метод удаления пользователя из группы, как себя, так и другого по id
func (p *PGStore) RemoveUserFromGroup(ctx context.Context, list *models.TodoList) error {
	q := `DELETE FROM user_groups WHERE user_id=$1 AND group_id=$2 `
	if _, err := p.client.Exec(ctx, q, list.AnotherUserID, list.GroupID); err != nil {
		// TODO ERROR
		return err
	}
	return nil
}

func (p *PGStore) GetAll(ctx context.Context, u *models.User) error {
	return nil
}

func (p *PGStore) UpdateUser(ctx context.Context, u *models.User) error {
	return nil
}

func (p *PGStore) UpdateList(ctx context.Context, list *models.User) error {
	return nil
}

// hashPassword - функция хэширования пароля методом sha256
func (p *PGStore) hashPassword(pass string) string {
	h := hmac.New(sha256.New, []byte("password"))
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}
