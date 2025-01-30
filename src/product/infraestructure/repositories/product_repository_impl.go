// repositories/product_repository_impl.go
package repositories

import (
	"database/sql"
	"demo/src/product/domain/entities"
	"fmt"
)

type ProductRepository interface {
    Create(product *entities.Product) error
    GetByID(id int) (*entities.Product, error)
    Update(product *entities.Product) error
    Delete(id int) error
}

type ProductRepositoryImpl struct {
    DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
    return &ProductRepositoryImpl{DB: db}
}

func (r *ProductRepositoryImpl) Create(product *entities.Product) error {
    query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
    _, err := r.DB.Exec(query, product.Name, product.Description, product.Price)
    return err
}

func (r *ProductRepositoryImpl) GetByID(id int) (*entities.Product, error) {
    query := "SELECT id, name, description, price FROM products WHERE id = ?"
    row := r.DB.QueryRow(query, id)
    var product entities.Product
    err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepositoryImpl) Update(product *entities.Product) error {
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

func (r *ProductRepositoryImpl) Delete(id int) error {
    query := "DELETE FROM products WHERE id=?"
    _, err := r.DB.Exec(query, id)
    return err
}