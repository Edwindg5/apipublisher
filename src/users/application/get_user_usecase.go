package application

import (
	"demo/src/users/domain/entities"
	"demo/src/users/infraestructure/repositories"
)

type GetUserUseCase struct {
	Repo *repositories.GetUserRepository
}

func NewGetUserUseCase(repo *repositories.GetUserRepository) *GetUserUseCase {
	return &GetUserUseCase{Repo: repo}
}

// Obtener todos los usuarios
func (uc *GetUserUseCase) GetAllUsers() ([]entities.User, error) {
	return uc.Repo.GetUsers()
}

// Obtener un usuario por ID
func (uc *GetUserUseCase) GetUserByID(id int) (*entities.User, error) {
	return uc.Repo.GetUserByID(id)
}
