package repository

import (
	"fmt"
	"time"

	"github.com/qubicfire/sl-service/internal/domain"
)

type MemoryRepository struct {
	emails map[string]*domain.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		emails: make(map[string]*domain.User),
	}
}

func (r *MemoryRepository) Close() error {
	clear(r.emails)
	return nil
}

func (r *MemoryRepository) CreateUser(name string, email string, password string) (*domain.User, error) {
	if _, ok := r.emails[email]; ok {
		return nil, fmt.Errorf("Failed to create user, because he already exist")
	}

	user := &domain.User{
		ID:        int64(len(r.emails)),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Local(),
	}
	r.emails[email] = user
	return user, nil
}

func (r *MemoryRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := r.emails[email]
	if user == nil {
		return user, fmt.Errorf("Failed to find user by email %s", email)
	}

	return user, nil
}

func (r *MemoryRepository) CreateLink(userID int64, url string, shortUrl string) (*domain.Link, error) {
	return nil, nil
}

func (r *MemoryRepository) DeleteLinkByID(userID int64, shortUrl string) error {
	return nil
}
