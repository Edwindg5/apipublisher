package application

import (
	"errors"
	"demo/src/users/domain/repositories"
	"demo/src/utils"
)

type LoginUserUseCase struct {
	UserRepo repositories.UserRepository
}

func NewLoginUserUseCase(repo repositories.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{UserRepo: repo}
}

func (uc *LoginUserUseCase) Execute(email, password string) (string, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil || !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return uc.GenerateToken(user.ID, user.Email)
}

func (uc *LoginUserUseCase) GenerateToken(userID int, email string) (string, error) {
	return utils.GenerateJWT(userID, email)
}
