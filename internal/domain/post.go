package domain

import (
	"time"
)

type Post struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	Title     string `db:"title" json:"title"`
	Text      string `db:"text" json:"text"`
	Images    []string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
