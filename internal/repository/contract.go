package repository

import (
	"GO-FJ/internal/domain"
	"context"
)

type PostRepository interface {
	Create(c context.Context, post *domain.Post) error
	GetByTitle(c context.Context, title string) ([]domain.Post, error)
	GetByID(c context.Context, id string) (domain.Post, error)
	GetByUserID(c context.Context, userID string) ([]domain.Post, error)
	UpdatePost(c context.Context, newPost domain.Post) error
	DeletePost(c context.Context, id string) error
}

type UserRepository interface {
	Create(c context.Context, user *domain.User) error
	GetByEmail(c context.Context, email string) (domain.User, error)
	GetByID(c context.Context, id string) (domain.User, error)
}

type CommentRepository interface {
	Create(c context.Context, comment *domain.Comment) error
	GetByID(c context.Context, id string) (domain.Comment, error)
	GetByPostID(c context.Context, postID string) ([]domain.Comment, error)
	UpdateComment(c context.Context, newComment domain.Comment) error
	DeleteComment(c context.Context, id string) error
}
