package transport

import (
	"context"
	"errors"
	"github.com/CyrilSbrodov/ToDoList/internal/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "safasGdf123fgdfg1SFa"
	tokenTTL   = 12 * time.Hour
)

// tokenClaims - структура токена
type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

// CreateUser - метод создания пользователя
func (t *Transport) CreateUser(ctx context.Context, u *models.User) (string, error) {
	return t.NewUser(ctx, u)
}

// GenerateToken - метод генерации нового токена авторизации
func (t *Transport) GenerateToken(ctx context.Context, u *models.User) (string, error) {
	id, err := t.Auth(ctx, u)
	if err != nil {
		t.logger.Error("error auth", err)
		return "", err
	}
	// создание нового токена с временем действия 12 часов. Время установлено в константе.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(signingKey))
}

// ParseToken - проверка токена при авторизации
func (t *Transport) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", errors.New("token claims are not type *tokenClaims or not Valid")
	}
}
