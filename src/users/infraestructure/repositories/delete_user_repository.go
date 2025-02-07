package repositories

import (
	"database/sql"
	"log"
)

type DeleteUserRepository struct {
	DB *sql.DB
}

func NewDeleteUserRepository(db *sql.DB) *DeleteUserRepository {
	return &DeleteUserRepository{DB: db}
}

// Eliminar usuario por ID
func (r *DeleteUserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar usuario:", err)
		return err
	}
	return nil
}
