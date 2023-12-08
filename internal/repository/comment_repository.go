package repository

import (
	"GO-FJ/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) Create(c context.Context, comment *domain.Comment) error {
	_, err := cr.db.ExecContext(
		c,
		queryCreateComment,
		comment.Message,
		comment.PostID,
		comment.UserID,
		comment.ParentID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		logrus.Errorf("cannot insert comment into comments: %s", err.Error())
	}

	return err
}

func (cr *commentRepository) GetByID(c context.Context, id string) (domain.Comment, error) {
	comment := domain.Comment{}
	err := cr.db.GetContext(c, &comment, queryGetCommentByID, id)
	if err != nil {
		return comment, nil
	}

	return comment, err
}

func (cr *commentRepository) GetByPostID(c context.Context, postID string) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := cr.db.SelectContext(
		c,
		&comments,
		queryGetCommentsByPostID,
		postID,
	)
	if err != nil {
		logrus.Errorf("cannot get comments with post_id %s: %s", postID, err.Error())
		return nil, nil
	}

	return comments, nil
}

func (cr *commentRepository) UpdateComment(c context.Context, newComment domain.Comment) error {
	_, err := cr.db.ExecContext(
		c,
		queryUpdateComment,
		newComment.Message,
		time.Now(),
		newComment.ID,
	)
	if err != nil {
		logrus.Errorf("cannot update comment with id %d: %s", newComment.ID, err.Error())
		return err
	}

	return nil
}

func (cr *commentRepository) DeleteComment(c context.Context, id string) error {
	_, err := cr.db.ExecContext(
		c,
		queryDeleteComment,
		id,
	)
	if err != nil {
		logrus.Errorf("cannot delete comment with id %s: %s", id, err.Error())
		return err
	}

	return nil
}
