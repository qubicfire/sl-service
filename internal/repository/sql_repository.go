package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/qubicfire/sl-service/internal/domain"
)

type SQLRepository struct {
	db *pgx.Conn
}

func NewDBRepository() (*SQLRepository, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Failed to load .env file")
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	timeout, err := strconv.Atoi(os.Getenv("CONNECTION_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to convert CONNECTION_TIMEOUT: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		db.Close(ctx)
		return nil, fmt.Errorf("Failed to ping db: %w", err)
	}

	slog.Info("Successfuly connected to db")
	return &SQLRepository{db: db}, nil
}

func (r *SQLRepository) Close() error {
	return r.db.Close(context.Background())
}

func (r *SQLRepository) CreateUser(name string, email string, password string) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id`
	var user domain.User
	err := r.db.QueryRow(
		context.Background(),
		query,
		name,
		email,
		password,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, name, email, password, created_at 
		FROM users 
		WHERE email = $1 
		LIMIT 1
	`
	var user domain.User

	err := r.db.QueryRow(context.Background(), query, email).Scan(
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

		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &user, nil
}

func (r *SQLRepository) CreateLink(userID int64, url string, shortUrl string) (*domain.Link, error) {
	query := `
		INSERT INTO links (url, short_url, user_id)
		VALUES ($1, $2, $3)
		RETURNING id`
	var link domain.Link
	err := r.db.QueryRow(
		context.Background(),
		query,
		url,
		shortUrl,
		userID,
	).Scan(&link.ID)

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *SQLRepository) DeleteLinkByID(userID int64, shortUrl string) error {
	query := `
		DELETE FROM links
		WHERE user_id = $1
		AND short_url = $2
	`
	var id int64
	err := r.db.QueryRow(
		context.Background(),
		query,
		userID,
		shortUrl,
	).Scan(id)
	return err
}
