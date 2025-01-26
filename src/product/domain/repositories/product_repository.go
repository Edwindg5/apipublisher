package repositories

import (
	"database/sql"
	"demo/src/product/domain/entities"
)

type ProductRepository struct {
	DB *sql.DB
}

func (repo *ProductRepository) Create(product entities.Product) error {
	query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, product.Name, product.Description, product.Price)
	return err
}

func (repo *ProductRepository) GetAll() ([]entities.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepository) Update(product entities.Product) error {
	query := "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	return err
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}
