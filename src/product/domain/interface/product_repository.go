package interfaces

import "demo/src/product/domain/entities"

// ProductRepository define el contrato que cualquier repositorio de productos debe cumplir.
type ProductRepository interface {
	Create(product *entities.Product) error
	GetByID(id int) (*entities.Product, error)
	GetAll() ([]*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id int) error
}
