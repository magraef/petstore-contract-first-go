// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlcgen

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	// Insert a new pet with an existing or new category
	// @name - Name of the pet
	// @categoryName - Name of the category
	// @categoryId - ID of an existing category (optional)
	CreatePet(ctx context.Context, db DBTX, arg CreatePetParams) (CreatePetRow, error)
	// Delete a pet by ID
	// :id - ID of the pet to delete
	DeletePet(ctx context.Context, db DBTX, id int32) (pgconn.CommandTag, error)
	// Find all pets optionally filtered by category.name
	// :categorieNames - Optional filter by category.name
	// :limit - Limit the number of results
	// :from - from which id to collect the results
	FindAllPets(ctx context.Context, db DBTX, arg FindAllPetsParams) ([]FindAllPetsRow, error)
	// Get a pet by ID
	// :id - ID of the pet
	GetPetById(ctx context.Context, db DBTX, id int32) (GetPetByIdRow, error)
	// Update an existing pet by ID
	// :id - ID of the pet to update
	// :name - New name of the pet
	// :category_name - New category ID of the pet
	UpdatePet(ctx context.Context, db DBTX, arg UpdatePetParams) (pgconn.CommandTag, error)
}

var _ Querier = (*Queries)(nil)
