package internal

import (
	"context"
)

type PetsRepository interface {
	Repository
	PetsResource
}

type Repository interface {
	Close()
	Ping(ctx context.Context) error
}
