package urlshortener

import (
	"time"
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

func (s *Shortener) Short(url string) (string, error) {
	return "", nil
}
