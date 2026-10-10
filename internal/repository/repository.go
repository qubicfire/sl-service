package repository

import (
	"github.com/qubicfire/sl-service/internal/domain"
)

type Repository interface {
	CreateUser(name string, email string, password string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
