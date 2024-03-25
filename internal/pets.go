package internal

import (
	"context"
)

type PetId int64

// Pet defines model for Pet.
type Pet struct {
	Id       *PetId
	Name     string
	Category Category
}

// FindAllPetsFilters defines filters for finding pets based on categories.
type FindAllPetsFilters struct {
	Categories *[]string
}

// PetsResource defines operations that can be performed on pets.
type PetsResource interface {
	// FindAllPets retrieves a list of pets based on provided filters and options.
	FindAllPets(ctx context.Context, filter FindAllPetsFilters, options Options) (*[]Pet, error)
	// CreatePet adds a new pet to the system.
	CreatePet(ctx context.Context, pet Pet) (Pet, error)
	// UpdatePet updates an existing pet with new information.
	UpdatePet(ctx context.Context, id PetId, newPet Pet) error
	// DeletePet removes a pet from the system based on its ID.
	DeletePet(ctx context.Context, id PetId) error
	// GetPetById retrieves a pet from the system based on its ID.
	GetPetById(ctx context.Context, id PetId) (Pet, error)
}

var _ PetsResource = (*PetsResourceImpl)(nil)

// PetsResourceImpl represents the implementation of the PetsResource interface.
// It encapsulates the business logic related to pets, independent of persistence and incoming adapter logic.
type PetsResourceImpl struct {
	repository PetsRepository
}

// NewPetResourceImpl creates a new instance of PetsResourceImpl with the provided repository.
func NewPetResourceImpl(repository PetsRepository) PetsResourceImpl {
	return PetsResourceImpl{repository: repository}
}

func (p PetsResourceImpl) FindAllPets(ctx context.Context, filter FindAllPetsFilters, options Options) (*[]Pet, error) {
	return p.repository.FindAllPets(ctx, filter, options)
}

func (p PetsResourceImpl) CreatePet(ctx context.Context, pet Pet) (Pet, error) {
	return p.repository.CreatePet(ctx, pet)
}

func (p PetsResourceImpl) UpdatePet(ctx context.Context, id PetId, newPet Pet) error {
	return p.repository.UpdatePet(ctx, id, newPet)
}

func (p PetsResourceImpl) DeletePet(ctx context.Context, id PetId) error {
	return p.repository.DeletePet(ctx, id)
}

func (p PetsResourceImpl) GetPetById(ctx context.Context, id PetId) (Pet, error) {
	return p.repository.GetPetById(ctx, id)
}
