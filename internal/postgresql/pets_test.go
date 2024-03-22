package postgresql

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	postgresContainer testcontainers.Container
	pp                *pgxpool.Pool
)

// TestMain runs before all tests
func TestMain(m *testing.M) {
	postgresContainer, connString, err := startPostgresContainer()

	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
		os.Exit(1)
	}

	pp = NewPgxPool(connString, "test")

	code := m.Run()

	defer pp.Close()
	// Start PostgreSQL container
	// Run tests
	defer postgresContainer.Stop(context.Background(), nil)
	os.Exit(code)
}

func startPostgresContainer() (testcontainers.Container, string, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432"),
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

	connString := fmt.Sprintf("postgres://postgres:postgres@%s:%s/postgres?sslmode=disable", ip, port.Port())

	return postgresContainer, connString, nil
}

func TestRepository_CreatePet(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
		pet internal.Pet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should create pet",
			args: args{
				ctx: context.Background(),
				pet: internal.Pet{
					Name: "Testpet",
					Category: internal.Category{
						Name: "fish",
					},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(pp)
			_, err := r.CreatePet(tt.args.ctx, tt.args.pet)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_DeletePet(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  internal.PetId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.DeletePet(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeletePet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_FindAllPets(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		filter  internal.FindAllPetsFilters
		options internal.Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]internal.Pet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.FindAllPets(tt.args.ctx, tt.args.filter, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAllPets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPets() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetPetById(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  internal.PetId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    internal.Pet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetPetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_UpdatePet(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		id     internal.PetId
		newPet internal.Pet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.UpdatePet(tt.args.ctx, tt.args.id, tt.args.newPet); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
