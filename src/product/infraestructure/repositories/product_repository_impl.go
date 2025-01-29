package repositories

import (
    "database/sql"
    "demo/src/product/domain/entities"
)

type ProductRepository interface {
    Create(product *entities.Product) error
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
