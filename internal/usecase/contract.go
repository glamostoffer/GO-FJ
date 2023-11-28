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
