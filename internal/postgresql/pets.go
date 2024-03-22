package postgresql

import (
	"context"

	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/postgresql/sqlcgen"
)

var _ internal.PetsResource = (*Repository)(nil)

func (r *Repository) FindAllPets(ctx context.Context, filter internal.FindAllPetsFilters, options internal.Options) (*[]internal.Pet, error) {

	pets, err := r.querier.FindAllPets(ctx, r.db, sqlcgen.FindAllPetsParams{
		Maxlimit:       int32(options.Limit),
		Startfrom:      int32(options.From),
		CategorieNames: *filter.Categories,
	})
	if err != nil {
		return nil, err
	}

	var result []internal.Pet
	for _, p := range pets {
		id := internal.PetId(p.ID)
		cid := internal.CategoryId(p.CategoryID)

		result = append(result, internal.Pet{
			Id:   &id,
			Name: p.Name,
			Category: internal.Category{
				Id:   &cid,
				Name: p.CategoryName,
			},
		})
	}

	return &result, nil
}

func (r *Repository) CreatePet(ctx context.Context, pet internal.Pet) (internal.Pet, error) {
	id, err := r.querier.CreatePet(ctx, r.db, sqlcgen.CreatePetParams{
		Name:         pet.Name,
		Categoryname: pet.Category.Name,
	})
	if err != nil {
		return internal.Pet{}, err
	}

	petId := internal.PetId(id.PetID)
	categoryId := internal.CategoryId(id.CategoryID)
	return internal.Pet{
		Id:   &petId,
		Name: id.PetName,
		Category: internal.Category{
			Id:   &categoryId,
			Name: id.CategoryName,
		},
	}, nil
}

func (r *Repository) UpdatePet(ctx context.Context, id internal.PetId, newPet internal.Pet) error {
	return r.querier.UpdatePet(ctx, r.db, sqlcgen.UpdatePetParams{
		Petid:        int32(id),
		Name:         newPet.Name,
		Categoryname: newPet.Category.Name,
	})
}

func (r *Repository) DeletePet(ctx context.Context, id internal.PetId) error {
	return r.querier.DeletePet(ctx, r.db, int32(id))
}

func (r *Repository) GetPetById(ctx context.Context, id internal.PetId) (internal.Pet, error) {

	petById, err := r.querier.GetPetById(ctx, r.db, int32(id))
	if err != nil {
		return internal.Pet{}, err
	}

	pid := internal.PetId(petById.ID)
	cid := internal.CategoryId(petById.CategoryID)

	return internal.Pet{
		Id:   &pid,
		Name: petById.Name,
		Category: internal.Category{
			Id:   &cid,
			Name: petById.CategoryName,
		},
	}, nil
}
