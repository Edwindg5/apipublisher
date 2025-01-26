package repositories

import (
	"database/sql"
	"errors"
	"demo/src/users/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	var user entities.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

