package repository

import (
	"GO-FJ/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	//var lastInsertId int64
	_, err := ur.db.ExecContext(
		c,
		queryCreateUser,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		logrus.Errorf("cannot insert user into users: %s", err.Error())
	}

	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	u := domain.User{}
	err := ur.db.GetContext(c, &u, queryGetUserByEmail, email)
	if err != nil {
		return u, nil
	}

	return u, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	u := domain.User{}
	err := ur.db.GetContext(c, &u, queryGetUserByID, id)
	if err != nil {
		return u, nil
	}

	return u, err
}
