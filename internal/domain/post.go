package domain

import (
	"time"
)

type Post struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"userID"`
	Title     string    `db:"title" json:"title"`
	Text      string    `db:"text" json:"text"`
	Images    []string  `json:"images"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
