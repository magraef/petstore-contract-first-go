//go:generate oapi-codegen --config ./.oapi-codegen/oapi-codegen-server.cfg.yaml ../../docs/openapi.yaml
//go:generate oapi-codegen --config ./.oapi-codegen/oapi-codegen-types.cfg.yaml ../../docs/openapi.yaml

package http

import (
	"context"

	"github.com/magraef/petstore-contract-first-go/internal"
)

var _ StrictServerInterface = (*PetstoreController)(nil)

type PetstoreController struct {
	resource internal.PetsResource
}

func NewPetstoreController(resource internal.PetsResource) PetstoreController {
	return PetstoreController{resource: resource}
}

func (p PetstoreController) GetPets(ctx context.Context, request GetPetsRequestObject) (GetPetsResponseObject, error) {

	var filters []string

	if request.Params.Category != nil {
		for _, filter := range *request.Params.Category {
			filters = append(filters, filter)
		}
	}

	pets, err := p.resource.FindAllPets(ctx,
		internal.FindAllPetsFilters{Categories: &filters},
		internal.NewOptions(internal.WithLimit(request.Params.Limit), internal.WithFrom(request.Params.From)),
	)

	if err != nil {
		return nil, err
	}

	result := GetPets200JSONResponse{}

	for _, p := range *pets {
		result = append(result, Pet{
			Id:   (*int64)(p.Id),
			Name: p.Name,
			Category: Category{
				Id:   (*int64)(p.Category.Id),
				Name: p.Category.Name,
			},
		})
	}

	return result, nil
}

func (p PetstoreController) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {

	pet, err := p.resource.CreatePet(ctx, internal.Pet{
		Name: request.Body.Name,
		Category: internal.Category{
			Name: request.Body.Category.Name,
		},
	})

	if err != nil {
		return nil, err
	}

	return AddPet201JSONResponse{Id: (*int64)(pet.Id),
		Name: pet.Name,
		Category: Category{
			Id:   (*int64)(pet.Category.Id),
			Name: pet.Category.Name,
		},
	}, nil
}

func (p PetstoreController) UpdatePet(ctx context.Context, request UpdatePetRequestObject) (UpdatePetResponseObject, error) {
	targetId := internal.PetId(request.PetId)
	cId := internal.CategoryId(*request.Body.Category.Id)

	if err := p.resource.UpdatePet(ctx, targetId, internal.Pet{
		Id:   &targetId,
		Name: request.Body.Name,
		Category: internal.Category{
			Id:   &cId,
			Name: request.Body.Category.Name,
		},
	}); err != nil {
		return nil, err
	}

	return UpdatePet202Response{}, nil
}

func (p PetstoreController) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	if err := p.resource.DeletePet(ctx, internal.PetId(request.PetId)); err != nil {
		return nil, err
	}

	return DeletePet204Response{}, nil
}

func (p PetstoreController) GetPetById(ctx context.Context, request GetPetByIdRequestObject) (GetPetByIdResponseObject, error) {
	pet, err := p.resource.GetPetById(ctx, internal.PetId(request.PetId))
	if err != nil {
		return nil, err
	}

	return GetPetById200JSONResponse{
		Id:   (*int64)(pet.Id),
		Name: pet.Name,
		Category: Category{
			Id:   (*int64)(pet.Category.Id),
			Name: pet.Category.Name,
		},
	}, nil
}
