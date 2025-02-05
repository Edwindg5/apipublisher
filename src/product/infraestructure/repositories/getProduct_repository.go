// src/product/infraestructure/repositories/getProduct_repository.go
package repositories

import (
	"database/sql"
	"demo/src/product/domain/entities"
)

type GetProductRepository interface {
	GetByID(id int) (*entities.Product, error)
    GetAll() ([]*entities.Product, error)
}



type GetProductRepositoryImpl struct {
	DB *sql.DB
}

func NewGetProductRepository(db *sql.DB) GetProductRepository {
	return &GetProductRepositoryImpl{DB: db}
}

func (r *GetProductRepositoryImpl) GetByID(id int) (*entities.Product, error) {
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

func (r *GetProductRepositoryImpl) GetAll() ([]*entities.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
