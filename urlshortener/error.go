package urlshortener

import "errors"

var (
	ErrorKeyGen      = errors.New("error generating key")
	ErrorKeyNotFound = errors.New("original URL not found")
	ErrExpiredKey    = errors.New("expired key")
)
