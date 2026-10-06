package domain

import (
	"context"
	"time"
)

type Link struct {
	ID        string
	Original  string
	UserID    int64
	CreatedAt time.Time
}

type LinkRepository interface {
	Create(ctx context.Context, link Link) error
	GetByID(ctx context.Context, id string) (Link, error)
}
