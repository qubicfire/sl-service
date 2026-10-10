package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/qubicfire/sl-service/internal/domain"
)

type SQLRepository struct {
	database sql.DB
}

func (r *SQLRepository) CreateUser(name string, email string, password string) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	var user domain.User
	err := r.database.QueryRow(
		query,
		name,
		email,
		password,
		time.Now().Local(),
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = ? LIMIT = 1"
	var user domain.User

	err := r.database.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Maybe login or password was wrong")
		}
	}

	return &user, nil
}
