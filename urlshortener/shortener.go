package urlshortener

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Shortener struct {
	urlshortener URLShortener
	cache        Cache
}

type URLShortener interface {
	Storer
	Finder
}

type Cache interface {
	Storer
	Finder
}

type Finder interface {
	FindByKey(context.Context, string) (ShortnedURL, error)
}

type Storer interface {
	Store(context.Context, ShortnedURL) (ShortnedURL, error)
}

func New(urlshortener URLShortener, cache Cache) *Shortener {
	return &Shortener{
		urlshortener: urlshortener,
		cache:        cache,
	}
}

func (s *Shortener) Short(ctx context.Context, url string) (string, error) {
	const (
		maxkeyLen     = 11
		daysToExpired = 7
	)

	var (
		keyLen    = 8
		key       string
		expiredAt = time.Now().Add(time.Hour * 24 * daysToExpired)
	)

	for {
		if keyLen > maxkeyLen {
			return "", ErrorKeyGen
		}

		key = KeyGen(keyLen)

		_, err := s.urlshortener.FindByKey(ctx, key)
		if errors.Is(err, ErrorKeyNotFound) {
			break
		}

		if err != nil {
			slog.Error(fmt.Sprintf("failed to query duplicate key: %v", err))
			return "", err
		}

		keyLen++
	}

	_, err := s.urlshortener.Store(ctx, ShortnedURL{
		Id:        uuid.NewString(),
		Key:       key,
		LongURL:   url,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return "", err
	}

	return key, nil
}

func (s *Shortener) GetURL(ctx context.Context, key string) (ShortnedURL, error) {
	cachedShortnedURL, err := s.cache.FindByKey(ctx, key)
	if cachedShortnedURL.Key != "" {
		return cachedShortnedURL, nil
	}

	if err != nil && !errors.Is(err, ErrorKeyNotFound) {
		return ShortnedURL{}, fmt.Errorf("failed to query cached key: %w", err)
	}

	shortnedURL, err := s.urlshortener.FindByKey(ctx, key)
	if errors.Is(err, ErrorKeyNotFound) {
		return ShortnedURL{}, ErrorKeyNotFound
	}

	if err != nil {
		return ShortnedURL{}, fmt.Errorf("failed to query key: %w", err)
	}

	if shortnedURL.IsExpired() {
		return ShortnedURL{}, ErrExpiredKey
	}

	_, err = s.cache.Store(ctx, shortnedURL)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to cache key: %v", err))
	}

	return shortnedURL, nil
}
