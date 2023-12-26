package urlshortener

import "time"

type ShortnedURL struct {
	Id        string    `json:"id"`
	Key       string    `json:"key"`
	LongURL   string    `json:"longUrl"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (s ShortnedURL) IsExpired() bool {
	return time.Now().After(s.ExpiredAt)
}
