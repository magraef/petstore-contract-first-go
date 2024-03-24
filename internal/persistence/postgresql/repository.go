//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/persistence/postgresql/sqlcgen"
	"github.com/rs/zerolog/log"
)

var _ internal.Repository = (*Repository)(nil)
var _ internal.ReadinessCheck = (*Repository)(nil)

type Repository struct {
	db      *pgxpool.Pool
	querier sqlcgen.Querier
}

func NewPgxPool(postgreUrl string, db string) *pgxpool.Pool {
	pgxConfig, err := pgxpool.ParseConfig(postgreUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config for connection db")
	}
	pgxConfig.ConnConfig.TLSConfig.InsecureSkipVerify = true
	pgxConfig.ConnConfig.Database = db

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create connection db")
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to ping postgresql")
	}

	return pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	repo := &Repository{db: pool, querier: sqlcgen.New()}
	err := repo.migrateSchema()
	if err != nil {
		log.Warn().Err(err).Msg("failed to run schema migration")
	}
	return repo
}

func (r *Repository) migrateSchema() error {
	return migrateSchema(r.db)
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) Check() error {
	return r.db.Ping(context.Background())
}
