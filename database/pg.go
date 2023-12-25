package database

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx"
)

var (
	db  *pgx.Conn
	err error
)

func StartPG() {
	db, err = pgx.Connect(pgx.ConnConfig{
		Host:     "localhost", // os.Getenv("BREVIORI_PG_HOST"),
		Port:     5432,
		User:     "postgres", // os.Getenv("BREVIORI_PG_USER"),
		Password: "password", // os.Getenv("BREVIORI_PG_PASSWORD"),
		Database: "postgres", // os.Getenv("BREVIORI_PG_DATABASE"),
	})
	if err != nil {
		slog.Error("failed connecting to database: %v", err)
		os.Exit(1)
	}

	slog.Info("connected to pg database")
}

func GetPGInstance() *pgx.Conn {
	return db
}
