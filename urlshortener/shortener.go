package urlshortener

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type Shortener struct {
	*deps
}

type ShortnedURL struct {
	Key       string
	LongURL   string
	ExpiredAt time.Time
}

func New() *Shortener {
	return &Shortener{
		deps: buildDependency(),
	}
}

func (s *Shortener) Short(ctx context.Context, url string) (string, error) {
	const (
		maxkeyLen         = 11
		checkDuplicateKey = `
			select id from url_key_pairs where key = $1
		`
		daysToExpired = 7
	)

	var (
		keepShortening = true
		keyLen         = 8
		key            string
		expiredAt      = time.Now().Add(time.Hour * 24 * daysToExpired)
		db             = s.deps.db
	)

	for keepShortening {
		if keyLen > maxkeyLen {
			return "", ErrorKeyGen
		}

		key = KeyGen(keyLen)
		var id string

		err := db.QueryRowEx(ctx, checkDuplicateKey, nil, key).Scan(&id)
		if errors.Is(err, pgx.ErrNoRows) {
			keepShortening = false
			continue
		}

		if err != nil {
			slog.Error(fmt.Sprintf("failed to query duplicate key: %v", err))
			return "", err
		}

		keyLen++
	}

	_, err := db.ExecEx(ctx, `
			insert into url_key_pairs (id, key, url, expired_at) values($1, $2, $3, $4)
		`, nil, uuid.NewString(), key, url, expiredAt)
	if err != nil {
		return "", err
	}

	return key, nil
}
