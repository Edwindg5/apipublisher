package repositories

import (
	"database/sql"
	"demo/src/users/domain/entities"
	"log"
)

type GetUserRepository struct {
	DB *sql.DB
}

func NewGetUserRepository(db *sql.DB) *GetUserRepository {
	return &GetUserRepository{DB: db}
}

// Obtener todos los usuarios
func (r *GetUserRepository) GetUsers() ([]entities.User, error) {
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

// Obtener usuario por ID
func (r *GetUserRepository) GetUserByID(id int) (*entities.User, error) {
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
