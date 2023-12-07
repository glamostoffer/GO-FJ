package repository

import (
	"GO-FJ/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type postRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) PostRepository {
	return &postRepository{db}
}

func (pr *postRepository) Create(c context.Context, post *domain.Post) error {
	tx, err := pr.db.Begin()
	if err != nil {
		tx.Rollback()
		logrus.Errorf("cannot start transaction: %s", err.Error())
		return err
	}

	var lastInsertId int64

	err = tx.QueryRowContext(
		c,
		queryCreatePost,
		post.Title,
		post.Text,
		post.CreatedAt,
		post.UpdatedAt,
		post.UserID,
	).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("cannot insert post into posts table: %s", err.Error())
		return err
	}

	var lastImageId int64
	for _, image := range post.Images {
		err = tx.QueryRowContext(
			c,
			queryCreateImage,
			image,
			lastInsertId,
		).Scan(&lastImageId)
		if err != nil {
			tx.Rollback()
			logrus.Errorf("cannot insert image into images table: %s", err.Error())
			return err
		}
	}

	tx.Commit()
	return nil
}

func (pr *postRepository) GetByTitle(c context.Context, title string) ([]domain.Post, error) {
	var posts []domain.Post
	err := pr.db.SelectContext(
		c,
		&posts,
		queryGetPostByTitle,
		title,
	)
	if err != nil {
		logrus.Errorf("cannot get posts with title %s: %s", title, err.Error())
		return nil, err
	}

	for _, post := range posts {
		var images []string
		err = pr.db.SelectContext(
			c,
			&images,
			queryGetImageByPostID,
			post.ID,
		)
		if err != nil {
			logrus.Errorf("cannot get images with post id %d: %s", post.ID, err.Error())
			//return nil, err
		}

		post.Images = images
	}

	return posts, nil
}

func (pr *postRepository) GetByID(c context.Context, id string) (domain.Post, error) {
	var post domain.Post
	err := pr.db.GetContext(
		c,
		&post,
		queryGetPostByID,
		id,
	)
	if err != nil {
		logrus.Errorf("cannot get post with id %s: %s", id, err.Error())
		return domain.Post{}, err
	}

	var images []string
	err = pr.db.SelectContext(
		c,
		&images,
		queryGetImageByPostID,
		id,
	)
	if err != nil {
		logrus.Errorf("cannot get images with post id %s: %s", id, err.Error())
		//return nil, err
	}

	post.Images = images

	return post, nil
}

func (pr *postRepository) GetByUserID(c context.Context, userID string) ([]domain.Post, error) {
	var posts []domain.Post
	err := pr.db.SelectContext(
		c,
		&posts,
		queryGetUserByUserID,
		userID,
	)
	if err != nil {
		logrus.Errorf("cannot get posts with title %s: %s", userID, err.Error())
		return nil, err
	}

	for _, post := range posts {
		var images []string
		err = pr.db.SelectContext(
			c,
			&images,
			queryGetImageByPostID,
			post.ID,
		)
		if err != nil {
			logrus.Errorf("cannot get images with post id %d: %s", post.ID, err.Error())
			//return nil, err
		}

		post.Images = images
	}

	return posts, nil
}
