package repository

import (
	"GO-FJ/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	var lastInsertId int64
	query := "INSERT INTO users(name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ur.db.QueryRowContext(c, query, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&lastInsertId)
	fmt.Println(lastInsertId)
	if err != nil {
		logrus.Errorf("cannot insert user into users: %s", err.Error())
	}

	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1"
	err := ur.db.QueryRowContext(c, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, nil
	}

	return u, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1"
	err := ur.db.QueryRowContext(c, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, nil
	}

	return u, err
}
