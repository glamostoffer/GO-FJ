package repository

import (
	"GO-FJ/internal/domain"
	"context"
)

type PostRepository interface {
	Create(c context.Context, post *domain.Post) error
	GetByTitle(c context.Context, title string) (domain.Post, error)
	GetByID(c context.Context, id string) (domain.Post, error)
	GetByUserID(c context.Context, userID string) (domain.Post, error)
}

type UserRepository interface {
	Create(c context.Context, user *domain.User) error
	GetByEmail(c context.Context, email string) (domain.User, error)
	GetByID(c context.Context, id string) (domain.User, error)
}
