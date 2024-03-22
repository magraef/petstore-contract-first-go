package internal

import "context"

type PetId int64

// Pet defines model for Pet.
type Pet struct {
	Id       *PetId
	Name     string
	Category Category
}

type FindAllPetsFilters struct {
	Categories *[]string
}

type PetsResource interface {
	FindAllPets(ctx context.Context, filter FindAllPetsFilters, options Options) (*[]Pet, error)
	CreatePet(ctx context.Context, pet Pet) (Pet, error)
	UpdatePet(ctx context.Context, id PetId, newPet Pet) error
	DeletePet(ctx context.Context, id PetId) error
	GetPetById(ctx context.Context, id PetId) (Pet, error)
}
