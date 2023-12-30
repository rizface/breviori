package database

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

var (
	dbConnPool *pgx.ConnPool
	err        error
)

func StartPG() {
	var port uint16

	if os.Getenv("BREVIORI_PG_PORT") == "" {
		port = 5432
	} else {
		intPort, err := strconv.ParseUint(os.Getenv("BREVIORI_PG_PORT"), 10, 16)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to parse BREVIORI_PG_PORT: %v", err))

			os.Exit(1)
		}

		port = uint16(intPort)
	}

	dbConnPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("BREVIORI_PG_HOST"),
			Port:     port,
			User:     os.Getenv("BREVIORI_PG_USER"),
			Password: os.Getenv("BREVIORI_PG_PASSWORD"),
			Database: os.Getenv("BREVIORI_PG_DATABASE"),
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
