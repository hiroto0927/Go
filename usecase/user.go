package usecase

import (
	"myapp/domain"
	"myapp/domain/repository"
)

type IUserUsecase interface {
	GetByUserID(ID string) (*domain.User, error)
	Insert(name, email string) error
}

type userUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) GetByUserID(ID string) (*domain.User, error) {
	user, err := uu.userRepository.GetByUserID(ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUsecase) Insert(name string, email string) error {
	err := uu.userRepository.Insert(name, email)

	if err != nil {
		return err
	}

	return nil
}
