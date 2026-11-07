package repository

import (
	"github.com/qubicfire/sl-service/internal/domain"
)

type Repository interface {
	CreateUser(name string, email string, password string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)

	CreateLink(userID int64, url string, shortUrl string) (*domain.Link, error)
	DeleteLinkByID(userID int64, shortUrl string) error

	Close() error
}
