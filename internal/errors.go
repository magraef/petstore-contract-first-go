package internal

import "errors"

var (
	ErrPetNotFound      = errors.New("pet not found")
	ErrPetAlreadyExists = errors.New("pet already exists")
)
