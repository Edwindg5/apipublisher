package repositories

import (
	"database/sql"
	"demo/src/product/domain/entities"
	"demo/src/product/domain/interface"
	"fmt"
)

type ProductsRepository struct {
	DB *sql.DB
}


func NewProductRepository(db *sql.DB) interfaces.ProductRepository {
	return &ProductsRepository{DB: db}
}


func (r *ProductsRepository) Create(product *entities.Product) error {
	query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price)
	return err
}


func (r *ProductsRepository) GetByID(id int) (*entities.Product, error) {
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

// Obtener todos los productos
func (r *ProductsRepository) GetAll() ([]*entities.Product, error) {
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

// Actualizar un producto
func (r *ProductsRepository) Update(product *entities.Product) error {
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

// Eliminar un producto
func (r *ProductsRepository) Delete(id int) error {
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
