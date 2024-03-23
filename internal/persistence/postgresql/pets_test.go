package postgresql

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pool *pgxpool.Pool
)

// TestMain runs before all tests
func TestMain(m *testing.M) {
	postgresContainer, connString, err := startPostgresContainer()

	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
		os.Exit(1)
	}

	pool = NewPgxPool(connString, "test")

	code := m.Run()

	defer pool.Close()
	// Start PostgreSQL container
	// Run tests
	defer postgresContainer.Stop(context.Background(), nil)
	os.Exit(code)
}

func startPostgresContainer() (testcontainers.Container, string, error) {
	env := make(map[string]string)
	env["POSTGRES_PASSWORD"] = "postgres"
	env["POSTGRES_DB"] = "test"

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432"),
		Env:          env,
	}

	postgresContainer, err := testcontainers.GenericContainer(context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		return nil, "", err
	}

	ip, err := postgresContainer.Host(context.Background())
	if err != nil {
		return nil, "", err
	}

	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, "", err
	}

	connString := fmt.Sprintf("postgres://postgres:postgres@%s:%s/test", ip, port.Port())

	return postgresContainer, connString, nil
}

func TestRepository_CreatePet(t *testing.T) {
	t.Run("create a new pet should be successful", func(t *testing.T) {
		r := NewRepository(pool)

		newPet := internal.Pet{
			Name: "CreatePet1",
			Category: internal.Category{
				Name: "CreatePet1",
			},
		}

		res, err := r.CreatePet(context.Background(), newPet)

		assert.NoError(t, err, "createPet failed")

		assert.NotNil(t, res, "pet should not be nil")
		assert.NotNil(t, res.Id, "pet ID should not be nil")
		assert.NotNil(t, res.Category, "pet category should not be nil")
		assert.NotNil(t, res.Category.Id, "pet category Id should not be nil")

		assert.NotNil(t, res.Category.Id, "pet category Id should not be nil")
		assert.NotNil(t, res.Category.Id, "pet category Id should not be nil")

		assert.EqualValues(t, newPet.Name, res.Name, "values should be equal")
		assert.EqualValues(t, newPet.Category.Name, res.Category.Name, "values should be equal")
	})

	t.Run("create a new pet with existing name should return ErrPetAlreadyExists", func(t *testing.T) {
		r := NewRepository(pool)

		newPet := internal.Pet{
			Name: "CreatePet2",
			Category: internal.Category{
				Name: "CreatePet2",
			},
		}

		res, err := r.CreatePet(context.Background(), newPet)
		assert.NoError(t, err, "createPet failed")
		assert.NotNil(t, res, "pet should not be nil")

		_, err = r.CreatePet(context.Background(), newPet)
		assert.ErrorIs(t, err, internal.ErrPetAlreadyExists, "createPet with existing name should return ErrPetAlreadyExists")
	})
}

func TestRepository_DeletePet(t *testing.T) {
	t.Run("delete existing pet should be successful", func(t *testing.T) {
		r := NewRepository(pool)

		pet, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "Test",
			Category: internal.Category{
				Name: "test2",
			},
		})

		assert.NoError(t, err, "error creating pet")
		assert.NotNil(t, pet, "pet should not be nil")

		err = r.DeletePet(context.Background(), *pet.Id)

		assert.NoError(t, err, "error deleting pet")
	})

	t.Run("delete not existing pet should return ErrPetNotFound", func(t *testing.T) {
		r := NewRepository(pool)

		err := r.DeletePet(context.Background(), 123)

		assert.ErrorIs(t, err, internal.ErrPetNotFound, "Result should be ErrPetNotFound")
	})
}

func TestRepository_FindAllPets(t *testing.T) {
	t.Run("find all pets with existing pets should return a list of pets", func(t *testing.T) {
		r := NewRepository(pool)

		_, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "allPets1",
			Category: internal.Category{
				Name: "AllPets1",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		got, err := r.FindAllPets(context.Background(), internal.FindAllPetsFilters{Categories: &[]string{}}, internal.NewOptions())
		assert.NoError(t, err, "Error finding pets")

		assert.NotNil(t, got, "Result should not be nil")
		assert.NotEmpty(t, *got, "Result should not be empty")
	})

}

func TestRepository_GetPetById(t *testing.T) {
	t.Run("get existing pet by id should return pet", func(t *testing.T) {
		r := NewRepository(pool)

		res, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "getExistingPetById",
			Category: internal.Category{
				Name: "getExistingPetById",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		got, err := r.GetPetById(context.Background(), *res.Id)
		assert.NoError(t, err, "Error finding pet by id")

		assert.EqualValues(t, res, got, "result should be the same")
	})

	t.Run("get none existing pet by id should return ErrPetNotFound", func(t *testing.T) {
		r := NewRepository(pool)

		_, err := r.GetPetById(context.Background(), internal.PetId(123))
		assert.ErrorIs(t, err, internal.ErrPetNotFound, "Error finding pet by id")
	})
}

func TestRepository_UpdatePet(t *testing.T) {
	t.Run("update existing pets with non existing categoryName should be successful", func(t *testing.T) {
		r := NewRepository(pool)

		res1, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithNonExistingCategoryName",
			Category: internal.Category{
				Name: "updatePetWithNonExistingCategoryName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		res2, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithNonExistingCategoryName2",
			Category: internal.Category{
				Name: "updatePetWithNonExistingCategoryName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		updatePet1 := internal.Pet{
			Name: res1.Name,
			Category: internal.Category{
				Name: res2.Category.Name,
			}}

		err = r.UpdatePet(context.Background(), *res1.Id, updatePet1)
		assert.NoError(t, err, "Error on update pet")

		result, err := r.GetPetById(context.Background(), *res1.Id)
		assert.NoError(t, err, "Error on update pet")

		assert.EqualValues(t, result, internal.Pet{
			Id:   res1.Id,
			Name: res1.Name,
			Category: internal.Category{
				Id:   res2.Category.Id,
				Name: updatePet1.Category.Name,
			},
		})
	})

	t.Run("update existing pet with existing categoryName should be successful", func(t *testing.T) {
		r := NewRepository(pool)

		res1, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingCategoryName",
			Category: internal.Category{
				Name: "updatePetWithExistingCategoryName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		res2, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingCategoryName2",
			Category: internal.Category{
				Name: "updatePetWithExistingCategoryName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		updatePet := internal.Pet{
			Name: "updatePetWithExistingCategoryName",
			Category: internal.Category{
				Name: res2.Category.Name,
			}}

		err = r.UpdatePet(context.Background(), *res1.Id, updatePet)
		assert.NoError(t, err, "Error on update pet")

		result, err := r.GetPetById(context.Background(), *res1.Id)
		assert.NoError(t, err, "Error on update pet")

		assert.EqualValues(t, result, internal.Pet{
			Id:   res1.Id,
			Name: res1.Name,
			Category: internal.Category{
				Id:   res2.Category.Id,
				Name: updatePet.Category.Name,
			},
		})
	})

	t.Run("update with non existing pet should return ErrPetNotFound", func(t *testing.T) {
		r := NewRepository(pool)

		err := r.UpdatePet(context.Background(), 123, internal.Pet{
			Name: "updateNonExistingPet",
			Category: internal.Category{
				Name: "updateNonExistingPet",
			},
		})
		assert.ErrorIs(t, err, internal.ErrPetNotFound, "should return ErrPetNotFound")
	})

	t.Run("update existing pet with existing name should return ErrPetAlreadyExists", func(t *testing.T) {
		r := NewRepository(pool)

		res1, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingName",
			Category: internal.Category{
				Name: "updatePetWithExistingName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		res2, err := r.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingName2",
			Category: internal.Category{
				Name: "updatePetWithExistingName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		err = r.UpdatePet(context.Background(), *res1.Id, internal.Pet{
			Name: res2.Name,
			Category: internal.Category{
				Name: res1.Category.Name,
			},
		})
		assert.ErrorIs(t, err, internal.ErrPetAlreadyExists, "Error on update pet")
	})
}
