package usecase

import (
	"GO-FJ/internal/domain"
	"context"
)

type SignupUsecase interface {
	Create(c context.Context, user *domain.User) error
	GetUserByEmail(c context.Context, email string) (domain.User, error)
	CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error)
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (domain.User, error)
	CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error)
}

type PostUsecase interface {
	Create(c context.Context, post *domain.Post) error
	GetByTitle(c context.Context, title string) ([]domain.Post, error)
	GetByID(c context.Context, id string) (domain.Post, error)
	GetByUserID(c context.Context, userID string) ([]domain.Post, error)
	UpdatePost(c context.Context, newPost domain.Post) error
	DeletePost(c context.Context, id string) error
}
