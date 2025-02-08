// src/users/infraestructure/repositories/repository_user.go

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

// Registrar usuario (POST)
func (r *UserRepository) RegisterUser(user entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error al registrar usuario:", err)
		return err
	}
	return nil
}

// Obtener todos los usuarios (GET)
func (r *UserRepository) GetUsers() ([]entities.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Obtener usuario por ID (GET)
func (r *UserRepository) GetUserByID(id int) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error al obtener usuario:", err)
		return nil, err
	}

	return &user, nil
}

// Actualizar usuario (PUT)
func (r *UserRepository) UpdateUser(user entities.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	res, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		log.Println("Error al actualizar usuario:", err)
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No se encontró el usuario con ID:", user.ID)
		return sql.ErrNoRows
	}

	return nil
}

// Eliminar usuario (DELETE)
func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar usuario:", err)
		return err
	}
	return nil
}
