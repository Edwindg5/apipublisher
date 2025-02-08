package interfaces

import "demo/src/users/domain/entities"

type IUserRepository interface {
	RegisterUser(user entities.User) error
	GetUsers() ([]entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id int) error
}
