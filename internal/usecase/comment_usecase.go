package usecase

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/repository"
	"context"
	"time"
)

type commentUsecase struct {
	commentRepository repository.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUsecase(repository repository.CommentRepository, timeout time.Duration) CommentUsecase {
	return &commentUsecase{
		commentRepository: repository,
		contextTimeout:    timeout,
	}
}

func (cu *commentUsecase) Create(c context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.Create(ctx, comment)
}

func (cu *commentUsecase) GetByID(c context.Context, id string) (domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.GetByID(ctx, id)
}

func (cu *commentUsecase) GetByPostID(c context.Context, postID string) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.GetByPostID(ctx, postID)
}

func (cu *commentUsecase) UpdateComment(c context.Context, newComment domain.Comment) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.UpdateComment(ctx, newComment)
}

func (cu *commentUsecase) DeleteComment(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.DeleteComment(ctx, id)
}
