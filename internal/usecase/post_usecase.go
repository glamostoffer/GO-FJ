package usecase

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/repository"
	"context"
	"time"
)

type postUsecase struct {
	postRepository repository.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(pr repository.PostRepository, timeout time.Duration) PostUsecase {
	return &postUsecase{
		postRepository: pr,
		contextTimeout: timeout,
	}
}

func (pu *postUsecase) Create(c context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.Create(ctx, post)
}

func (pu *postUsecase) GetByTitle(c context.Context, title string) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetByTitle(ctx, title)
}

func (pu *postUsecase) GetByID(c context.Context, id string) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetByID(ctx, id)
}

func (pu *postUsecase) GetByUserID(c context.Context, userID string) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetByUserID(ctx, userID)
}

func (pu *postUsecase) UpdatePost(c context.Context, id int64) (domain.Post, error) {
	return domain.Post{}, nil
}

func (pu *postUsecase) DeletePost(c context.Context, id int64) error {
	return nil
}
