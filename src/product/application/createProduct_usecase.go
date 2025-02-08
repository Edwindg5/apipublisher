package application

import (
	"demo/src/product/domain/entities"
	"demo/src/product/domain/interface"
)

type CreateProductsUsecase struct {
	Repo interfaces.ProductRepository
}

func NewCreateProductsUsecase(repo interfaces.ProductRepository) *CreateProductsUsecase {
	return &CreateProductsUsecase{Repo: repo}
}

func (uc *CreateProductsUsecase) CreateProduct(product *entities.Product) error {
	return uc.Repo.Create(product)
}
