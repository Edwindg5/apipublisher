package infrastructure

import (
	"errors"
	"demo/src/users/domain/entities"
	"demo/src/users/domain/repositories"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

// CreateUser creates a new user in the database
func (r *UserRepositoryImpl) CreateUser(user *entities.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByEmail retrieves a user by their email
func (r *UserRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
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
