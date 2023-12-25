package urlshortener

import (
	"github.com/jackc/pgx"
	"github.com/redis/go-redis/v9"
	"github.com/rizface/breviori/database"
)

type deps struct {
	db  *pgx.Conn
	rdb *redis.Client
}

func buildDependency() *deps {
	return &deps{
		db:  database.GetPGInstance(),
		rdb: database.GetRedisInstance(),
	}
}
