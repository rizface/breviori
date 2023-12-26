package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rizface/breviori/urlshortener"
)

type CacheShortnedURL struct {
	rdb *redis.Client
}

func NewRedisClient() *CacheShortnedURL {
	return &CacheShortnedURL{
		rdb: GetRedisInstance(),
	}
}

func (c *CacheShortnedURL) FindByKey(ctx context.Context, key string) (urlshortener.ShortnedURL, error) {
	shurl := urlshortener.ShortnedURL{}

	result, err := c.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return shurl, urlshortener.ErrorKeyNotFound
	}

	if err != nil {
		return shurl, err
	}

	if err := json.Unmarshal([]byte(result), &shurl); err != nil {
		return shurl, fmt.Errorf(fmt.Sprintf("failed to unmarshal cached key %s: %v", key, err))
	}

	return shurl, nil
}

func (c *CacheShortnedURL) Store(ctx context.Context, shortnedURL urlshortener.ShortnedURL) (urlshortener.ShortnedURL, error) {
	serialized, err := json.Marshal(shortnedURL)
	if err != nil {
		return shortnedURL, fmt.Errorf(fmt.Sprintf("failed to marshal shortnedURL: %v", err))
	}

	_, err = c.rdb.SetEx(ctx, shortnedURL.Key, serialized, time.Hour*24).Result()
	if err != nil {
		return shortnedURL, fmt.Errorf(fmt.Sprintf("failed to store shortnedURL: %v", err))
	}

	return shortnedURL, nil
}
