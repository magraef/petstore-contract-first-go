package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_schema "github.com/magraef/petstore-contract-first-go/internal/persistence/postgresql/.schema"
	"github.com/rs/zerolog/log"
	"github.com/z0ne-dev/mgx/v2"
)

func migrateSchema(pool *pgxpool.Pool) error {
	migrator, _ := mgx.New(
		mgx.Migrations(
			mgx.NewMigration("migrate required schemas for repository", func(ctx context.Context, commands mgx.Commands) error {
				file, err := _schema.Schema.ReadFile("schema.sql")
				if err != nil {
					return err
				}
				_, err = commands.Exec(ctx, string(file))
				if err != nil {
					return err
				}
				return nil
			}),
		),
		mgx.Log(logAdapter{}))

	if err := migrator.Migrate(context.Background(), pool); err != nil {
		return err
	}

	return nil
}

var _ mgx.Logger = (*logAdapter)(nil)

type logAdapter struct{}

func (l logAdapter) Log(msg string, data map[string]any) {

	var formattedData []interface{}
	for key, value := range data {
		formattedData = append(formattedData, fmt.Sprintf("%s: %v", key, value))
	}

	log.Info().Msgf("%s %s", msg, formattedData)
}
