package postgresql

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magraef/petstore-contract-first-go/internal"
)

var (
	pool *pgxpool.Pool
	repo *Repository
)

func TestPetsRepository_CreatePet(t *testing.T) {
	t.Run("create a new pet should be successful", func(t *testing.T) {

		// given
		newPet := internal.Pet{
			Name: "CreatePet",
			Category: internal.Category{
				Name: "CreatePet",
			},
		}

		// when
		res, err := repo.CreatePet(context.Background(), newPet)

		//then
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
		//given
		newPet := internal.Pet{
			Name: "CreatePetWithExistingName",
			Category: internal.Category{
				Name: "CreatePetWithExistingName",
			},
		}

		res, err := repo.CreatePet(context.Background(), newPet)
		assert.NoError(t, err, "createPet failed")
		assert.NotNil(t, res, "pet should not be nil")

		//when
		_, err = repo.CreatePet(context.Background(), newPet)

		//then
		assert.ErrorIs(t, err, internal.ErrPetAlreadyExists, "createPet with existing name should return ErrPetAlreadyExists")
	})
}

func TestPetsRepository_DeletePet(t *testing.T) {
	t.Run("delete existing pet should be successful", func(t *testing.T) {
		// given
		pet, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "Test",
			Category: internal.Category{
				Name: "test2",
			},
		})

		assert.NoError(t, err, "error creating pet")
		assert.NotNil(t, pet, "pet should not be nil")

		// when
		err = repo.DeletePet(context.Background(), *pet.Id)

		//then
		assert.NoError(t, err, "error deleting pet")
	})

	t.Run("delete not existing pet should return ErrPetNotFound", func(t *testing.T) {
		//given
		//when
		err := repo.DeletePet(context.Background(), 123)

		//then
		assert.ErrorIs(t, err, internal.ErrPetNotFound, "Result should be ErrPetNotFound")
	})
}

func TestPetsRepository_FindAllPets(t *testing.T) {
	t.Run("find all pets with existing pets should return a list of pets", func(t *testing.T) {

		//given
		_, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "allPets1",
			Category: internal.Category{
				Name: "AllPets1",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		//when
		got, err := repo.FindAllPets(context.Background(), internal.FindAllPetsFilters{Categories: &[]string{}}, internal.NewOptions())

		//then
		assert.NoError(t, err, "Error finding pets")

		assert.NotNil(t, got, "Result should not be nil")
		assert.NotEmpty(t, *got, "Result should not be empty")
	})

	t.Run("find all pets with non existing pets should return a empty list", func(t *testing.T) {
		//given
		_, err := pool.Exec(context.Background(), "TRUNCATE TABLE pets;")
		assert.NoError(t, err, "Error delete complete data pets")

		//when
		got, err := repo.FindAllPets(context.Background(), internal.FindAllPetsFilters{Categories: &[]string{}}, internal.NewOptions())

		//then
		assert.NoError(t, err, "Error finding pets")

		assert.Empty(t, *got, "Result should not be empty")
	})
}

func TestPetsRepository_GetPetById(t *testing.T) {
	t.Run("get existing pet by id should return pet", func(t *testing.T) {

		// given
		res, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "getExistingPetById",
			Category: internal.Category{
				Name: "getExistingPetById",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		//when
		got, err := repo.GetPetById(context.Background(), *res.Id)

		//then
		assert.NoError(t, err, "Error finding pet by id")

		assert.EqualValues(t, res, got, "result should be the same")
	})

	t.Run("get none existing pet by id should return ErrPetNotFound", func(t *testing.T) {
		// given
		// when
		_, err := repo.GetPetById(context.Background(), internal.PetId(123))

		// then
		assert.ErrorIs(t, err, internal.ErrPetNotFound, "Error finding pet by id")
	})
}

func TestPetsRepository_UpdatePet(t *testing.T) {
	t.Run("update existing pets with non existing categoryName should be successful", func(t *testing.T) {
		// given
		pet1, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithNonExistingCategoryName",
			Category: internal.Category{
				Name: "updatePetWithNonExistingCategoryName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		pet2, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithNonExistingCategoryName2",
			Category: internal.Category{
				Name: "updatePetWithNonExistingCategoryName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		// when
		petUpdate := internal.Pet{
			Name: pet1.Name,
			Category: internal.Category{
				Name: pet2.Category.Name,
			}}

		err = repo.UpdatePet(context.Background(), *pet1.Id, petUpdate)
		assert.NoError(t, err, "Error on update pet")

		// then
		result, err := repo.GetPetById(context.Background(), *pet1.Id)
		assert.NoError(t, err, "Error on update pet")

		assert.EqualValues(t, result, internal.Pet{
			Id:   pet1.Id,
			Name: pet1.Name,
			Category: internal.Category{
				Id:   pet2.Category.Id,
				Name: petUpdate.Category.Name,
			},
		})
	})

	t.Run("update existing pet with existing categoryName should be successful", func(t *testing.T) {
		// given
		pet1, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingCategoryName",
			Category: internal.Category{
				Name: "updatePetWithExistingCategoryName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		pet2, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingCategoryName2",
			Category: internal.Category{
				Name: "updatePetWithExistingCategoryName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		// when
		petUpdate := internal.Pet{
			Name: "updatePetWithExistingCategoryName",
			Category: internal.Category{
				Name: pet2.Category.Name,
			}}

		err = repo.UpdatePet(context.Background(), *pet1.Id, petUpdate)
		assert.NoError(t, err, "Error on update pet")

		// then
		result, err := repo.GetPetById(context.Background(), *pet1.Id)
		assert.NoError(t, err, "Error on update pet")

		assert.EqualValues(t, result, internal.Pet{
			Id:   pet1.Id,
			Name: pet1.Name,
			Category: internal.Category{
				Id:   pet2.Category.Id,
				Name: petUpdate.Category.Name,
			},
		})
	})

	t.Run("update with non existing pet should return ErrPetNotFound", func(t *testing.T) {
		// given
		// when
		err := repo.UpdatePet(context.Background(), 123, internal.Pet{
			Name: "updateNonExistingPet",
			Category: internal.Category{
				Name: "updateNonExistingPet",
			},
		})

		// then
		assert.ErrorIs(t, err, internal.ErrPetNotFound, "should return ErrPetNotFound")
	})

	t.Run("update existing pet with existing name should return ErrPetAlreadyExists", func(t *testing.T) {
		// given
		pet1, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingName",
			Category: internal.Category{
				Name: "updatePetWithExistingName",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		pet2, err := repo.CreatePet(context.Background(), internal.Pet{
			Name: "updatePetWithExistingName2",
			Category: internal.Category{
				Name: "updatePetWithExistingName2",
			},
		})
		assert.NoError(t, err, "Error on create pet")

		// when
		err = repo.UpdatePet(context.Background(), *pet1.Id, internal.Pet{
			Name: pet2.Name,
			Category: internal.Category{
				Name: pet1.Category.Name,
			},
		})

		// then
		assert.ErrorIs(t, err, internal.ErrPetAlreadyExists, "Error on update pet")
	})
}

// TestMain runs before all tests
func TestMain(m *testing.M) {
	postgresContainer, connString, err := startPostgresContainer()
	defer postgresContainer.Stop(context.Background(), nil)

	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
		os.Exit(1)
	}

	pool = NewPgxPool(connString, "test")
	defer pool.Close()

	repo = NewRepository(pool)

	code := m.Run()

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
