package repository

import (
	"GO-FJ/internal/domain"
	"context"
	"database/sql"
	"time"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	var lastInsertId int
	query := "INSERT INTO users(name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	err := ur.db.QueryRowContext(c, query, user.Name, user.Email, user.Password, "user", time.Now(), time.Now()).Scan(&lastInsertId)

	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1"
	err := ur.db.QueryRowContext(c, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return domain.User{}, nil
	}

	return u, nil
}
func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1"
	err := ur.db.QueryRowContext(c, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return domain.User{}, nil
	}

	return u, nil
}
