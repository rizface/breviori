package urlshortener

import "time"

type ShortnedURL struct {
	Key       string
	LongURL   string
	ExpiredAt time.Time
}

func (s ShortnedURL) IsExpired() bool {
	return time.Now().After(s.ExpiredAt)
}
