package domain

import (
	"time"
)

type Link struct {
	ID        string
	Url       string
	ShortUrl  string
	UserID    int64
	CreatedAt time.Time
}
