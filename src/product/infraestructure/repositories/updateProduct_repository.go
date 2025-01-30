package repositories

import (
    "database/sql"
    "demo/src/product/domain/entities"
    "fmt"
)

type UpdateProductRepository interface {
    Update(product *entities.Product) error
    GetByID(id int) (*entities.Product, error)
}

type UpdateProductRepositoryImpl struct {
    DB *sql.DB
}

func NewUpdateProductRepository(db *sql.DB) UpdateProductRepository {
    return &UpdateProductRepositoryImpl{DB: db}
}

func (r *UpdateProductRepositoryImpl) GetByID(id int) (*entities.Product, error) {
    var product entities.Product
    query := "SELECT id, name, description, price FROM products WHERE id = ?"
    err := r.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
    if err != nil {
        return nil, err
    }
    return &product, nil
}


func (r *UpdateProductRepositoryImpl) Update(product *entities.Product) error {
    query := "UPDATE products SET name=?, description=?, price=? WHERE id=?"
    result, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró el producto con ID %d", product.ID)
    }

    return nil
}
