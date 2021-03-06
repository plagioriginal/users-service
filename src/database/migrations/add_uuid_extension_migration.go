package migrations

import (
	"context"
	"database/sql"
	"log"
)

// Adds the uuid extension to the DB
func AddUUIDExtension(ctx context.Context, db *sql.DB, logger *log.Logger) error {
	query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

	_, err := db.ExecContext(ctx, query)

	return err
}

// Creates the new migration
func NewAddUuidExtensionMigration() Migration {
	return Migration{
		Name: "add-uuid-extension",
		Up:   AddUUIDExtension,
	}
}
