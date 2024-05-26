package models

import "errors"

var (
	ErrorUserConflict = errors.New("user or email already exists")
)
