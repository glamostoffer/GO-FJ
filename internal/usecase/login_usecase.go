package usecase

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/repository"
	"GO-FJ/util"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository repository.UserRepository, timeout time.Duration) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}
