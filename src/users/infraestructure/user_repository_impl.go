package infrastructure

import (
	"demo/src/users/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (repo *UserRepositoryImpl) CreateUser(user *entities.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return errors.New("error al guardar el usuario en la base de datos")
	}
	return nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Usuario no encontrado
		}
		return nil, errors.New("error al buscar el usuario en la base de datos")
	}
	return &user, nil
}


// GetUserByID retrieves a user by their ID
func (r *UserRepositoryImpl) GetUserByID(userID int) (*entities.User, error) {
	var user entities.User
	result := r.DB.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

// UpdateUser updates an existing user
func (r *UserRepositoryImpl) UpdateUser(user *entities.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user by their ID
func (r *UserRepositoryImpl) DeleteUser(userID int) error {
	result := r.DB.Delete(&entities.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
