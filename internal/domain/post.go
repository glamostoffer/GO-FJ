package domain

import (
	"time"
)

type Post struct {
	ID        int64
	UserID    int64
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
