package repository

import (
	"myapp/domain"
)

type IUserRepository interface {
	Insert(name, email string) error
	GetByUserID(ID string) (*domain.User, error)
}
