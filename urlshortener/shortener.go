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

func New() *Shortener {
	return &Shortener{
		deps: buildDependency(),
	}
}

func (s *Shortener) Short(ctx context.Context, url string) (string, error) {
	const (
		maxkeyLen     = 11
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

		err := db.QueryRowEx(ctx, `select id from url_key_pairs where key = $1`, nil, key).Scan(&id)
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

func (s *Shortener) GetURL(ctx context.Context, key string) (ShortnedURL, error) {
	var (
		db = s.deps.db
		kp ShortnedURL
	)

	err := db.QueryRowEx(ctx, `
		select id, url, key, expired_at from url_key_pairs where key = $1
	`, nil, key).
		Scan(&kp.Key, &kp.LongURL, &kp.Key, &kp.ExpiredAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return kp, ErrorKeyNotFound
	}

	if err != nil {
		return kp, fmt.Errorf("failed to query key: %w", err)
	}

	if kp.IsExpired() {
		return kp, ErrExpiredKey
	}

	return kp, nil
}
