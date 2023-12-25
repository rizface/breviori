package urlshortener

import (
	"github.com/jackc/pgx"
	"github.com/rizface/breviori/database"
)

type deps struct {
	db *pgx.Conn
}

func buildDependency() *deps {
	return &deps{
		db: database.GetConnectionInstance(),
	}
}
