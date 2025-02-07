package repositories

import (
	"database/sql"
	"demo/src/users/domain/entities"
	"log"
)

type PutUserRepository struct {
	DB *sql.DB
}

func NewPutUserRepository(db *sql.DB) *PutUserRepository {
	return &PutUserRepository{DB: db}
}

// Actualizar usuario por ID
func (r *PutUserRepository) UpdateUser(user entities.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	res, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		log.Println("Error al actualizar usuario:", err)
		return err
	}

	// Verificar si se actualizó algún registro
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No se encontró el usuario con ID:", user.ID)
		return sql.ErrNoRows
	}

	return nil
}
