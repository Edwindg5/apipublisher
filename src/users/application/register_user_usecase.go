package application

import (
	"demo/src/users/domain/entities"
	"demo/src/users/infraestructure/repositories"
)

type RegisterUserUseCase struct {
	Repo *repositories.UserRepository
}

func NewRegisterUserUseCase(repo *repositories.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{Repo: repo}
}

func (uc *RegisterUserUseCase) Execute(user entities.User) error {
	return uc.Repo.RegisterUser(user)
}
