package application

import (
	"demo/src/users/domain/entities"
	"demo/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
	UserRepo repositories.UserRepository
}

func NewRegisterUserUseCase(repo repositories.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{UserRepo: repo}
}

func (uc *RegisterUserUseCase) Execute(user *entities.User) error {
	return uc.UserRepo.CreateUser(user)
}
