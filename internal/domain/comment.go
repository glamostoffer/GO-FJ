package domain

import "time"

type Comment struct {
	ID        int64     `db:"id" json:"id"`
	Message   string    `db:"message" json:"message"`
	PostID    int64     `db:"post_id" json:"postID"`
	UserID    int64     `db:"user_id" json:"userID"`
	ParentID  int64     `db:"parent_id" json:"parentID"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
