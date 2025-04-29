package migrations

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	m *migrate.Migrate
}

func NewMigrator(db *sqlx.DB, migrationsPath string) (*Migrator, error) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("create migrate instance: %w", err)
	}

	return &Migrator{m: m}, nil
}

func (m *Migrator) Up() error {
	err := m.m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	log.Println("Migrations applied successfully.")
	return nil
}

func (m *Migrator) Down() error {
	err := m.m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate down: %w", err)
	}
	log.Println("Migrations rolled back successfully.")
	return nil
}

func (m *Migrator) SetVersion(version uint) error {
	if err := m.m.Migrate(version); err != nil {
		return fmt.Errorf("goto version %d: %w", version, err)
	}
	log.Printf("Migrated to version %d.\n", version)
	return nil
}
