package application

import "demo/src/users/infraestructure/repositories"

type DeleteUserUseCase struct {
	Repo *repositories.DeleteUserRepository
}

func NewDeleteUserUseCase(repo *repositories.DeleteUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{Repo: repo}
}

// Eliminar usuario
func (uc *DeleteUserUseCase) DeleteUser(id int) error {
	return uc.Repo.DeleteUser(id)
}
