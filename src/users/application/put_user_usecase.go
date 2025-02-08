package application

import (
	"demo/src/users/domain/entities"
	"demo/src/users/infraestructure/repositories"
)

type PutUserUseCase struct {
	Repo *repositories.UserRepository
}

func NewPutUserUseCase(repo *repositories.UserRepository) *PutUserUseCase {
	return &PutUserUseCase{Repo: repo}
}

// Actualizar usuario
func (uc *PutUserUseCase) UpdateUser(user entities.User) error {
	return uc.Repo.UpdateUser(user)
}
