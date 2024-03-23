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

var _ PetsResource = (*PetsResourceImpl)(nil)

type PetsResourceImpl struct {
	repository PetsRepository
}

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
