package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx"
)

var (
	dbConnPool *pgx.ConnPool
	err        error
)

func StartPG() {
	dbConnPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost", // os.Getenv("BREVIORI_PG_HOST"),
			Port:     5432,
			User:     "postgres", // os.Getenv("BREVIORI_PG_USER"),
			Password: "password", // os.Getenv("BREVIORI_PG_PASSWORD"),
			Database: "postgres", // os.Getenv("BREVIORI_PG_DATABASE"),
		},
		AfterConnect: func(c *pgx.Conn) error {
			fmt.Println("Acquire New Connection")
			return nil
		},
		MaxConnections: 10,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("failed connecting to database: %v", err))
		os.Exit(1)
	}

	slog.Info("connected to pg database")
}

func GetPGPool() *pgx.ConnPool {
	return dbConnPool
}
