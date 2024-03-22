package usecase

import (
	"context"

	"github.com/magraef/petstore-contract-first-go/internal"
)

var _ internal.PetsResource = (*PetUseCase)(nil)

type PetUseCase struct {
	repository internal.PetsRepository
}

func NewPetUseCase(repository internal.PetsRepository) PetUseCase {
	return PetUseCase{repository: repository}
}

func (p PetUseCase) FindAllPets(ctx context.Context, filter internal.FindAllPetsFilters, options internal.Options) (*[]internal.Pet, error) {
	return p.repository.FindAllPets(ctx, filter, options)
}

func (p PetUseCase) CreatePet(ctx context.Context, pet internal.Pet) (internal.Pet, error) {
	return p.repository.CreatePet(ctx, pet)
}

func (p PetUseCase) UpdatePet(ctx context.Context, id internal.PetId, newPet internal.Pet) error {
	return p.repository.UpdatePet(ctx, id, newPet)
}

func (p PetUseCase) DeletePet(ctx context.Context, id internal.PetId) error {
	return p.repository.DeletePet(ctx, id)
}

func (p PetUseCase) GetPetById(ctx context.Context, id internal.PetId) (internal.Pet, error) {
	return p.repository.GetPetById(ctx, id)
}
