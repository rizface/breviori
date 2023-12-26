package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/rizface/breviori/urlshortener"
)

type ShortnedURL struct {
	db *pgx.ConnPool
}

func NewShortnedURL() *ShortnedURL {
	return &ShortnedURL{
		db: GetPGPool(),
	}
}

func (s *ShortnedURL) FindByKey(ctx context.Context, key string) (urlshortener.ShortnedURL, error) {
	var shortnedURL urlshortener.ShortnedURL

	err := s.db.QueryRowEx(ctx, `select id, key, url, expired_at from url_key_pairs where key = $1`, nil, key).
		Scan(&shortnedURL.Id, &shortnedURL.Key, &shortnedURL.LongURL, &shortnedURL.ExpiredAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return shortnedURL, urlshortener.ErrorKeyNotFound
	}
	if err != nil {
		return shortnedURL, err
	}

	return shortnedURL, nil
}

func (s *ShortnedURL) Store(ctx context.Context, shortnedURL urlshortener.ShortnedURL) (urlshortener.ShortnedURL, error) {
	_, err := s.db.ExecEx(ctx, `insert into url_key_pairs (id, key, url, expired_at) values ($1, $2, $3, $4)`, nil,
		shortnedURL.Id, shortnedURL.Key, shortnedURL.LongURL, shortnedURL.ExpiredAt)
	if err != nil {
		return shortnedURL, err
	}

	return shortnedURL, nil
}
