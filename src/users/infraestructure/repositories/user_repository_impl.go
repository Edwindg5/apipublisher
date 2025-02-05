package repositories

import (
	"database/sql"
	"demo/src/users/domain/entities"

	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) RegisterUser(user entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error al registrar usuario:", err)
		return err
	}
	return nil
}
