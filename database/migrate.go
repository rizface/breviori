package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/stdlib"
)

func Migrate() {
	db := stdlib.OpenDBFromPool(dbConnPool)

	defer func() {
		if err := db.Close(); err != nil {
			slog.Warn(fmt.Sprintf("failed close connection for database migration: %v", err))
		}
	}()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create postgres driver for database migration: %v", err))
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migration", os.Getenv("BREVIORI_PG_DATABASE"), driver)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create migration instance: %v", err))
		os.Exit(1)
	}

	slog.Info("success migrate migrations")

	m.Up()
}
