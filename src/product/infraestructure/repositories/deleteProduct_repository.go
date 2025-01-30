// src/product/infraestructure/repositories/deleteProduct_repository.go
package repositories

import (
    "database/sql"
)

type DeleteProductRepository interface {
    Delete(id int) error
}


type DeleteProductRepositoryImpl struct {
    DB *sql.DB
}

func NewDeleteProductRepository(db *sql.DB) DeleteProductRepository {
    return &DeleteProductRepositoryImpl{DB: db}
}


func (r *DeleteProductRepositoryImpl) Delete(id int) error {
    query := "DELETE FROM products WHERE id = ?"
    result, err := r.DB.Exec(query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return sql.ErrNoRows
    }
    
    return nil
}
