package repo

import (
	"context"
	"errors"

	"github.com/densmart/sso-auth/internal/adapters/db/mockdb"
	"github.com/densmart/sso-auth/internal/adapters/db/postgres"
)

type Repo interface {
	Close()
	MigrationUp() error
	MigrationDown() error
	Roles
}

func NewRepo(ctx context.Context, DBType string) (Repo, error) {
	var database Repo
	switch DBType {
	case "postgres":
		pg, err := postgres.NewPostgresDB(ctx)
		if err != nil {
			return nil, err
		}
		database = pg
	case "mockdb":
		database = mockdb.NewMockDB()
	default:
		return nil, errors.New("invalid DB type")
	}

	return database, nil
}
