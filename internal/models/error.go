package models

import "errors"

// ошибки
var (
	ErrorUserConflict = errors.New("user or email already exists")
	ErrorUserNotFound = errors.New("user not found")
)
