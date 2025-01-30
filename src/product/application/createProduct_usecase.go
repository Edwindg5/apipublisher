//src/product/application/createProduct_usecase.go
package application

import (
    "demo/src/product/domain/entities"
    "demo/src/product/infraestructure/repositories"
)

type CreateProductsUsecase struct {
    Repo repositories.ProductRepository
}

func NewCreateProductsUsecase(repo repositories.ProductRepository) *CreateProductsUsecase {
    return &CreateProductsUsecase{Repo: repo}
}

func (uc *CreateProductsUsecase) CreateProduct(product *entities.Product) error {
    return uc.Repo.Create(product)
}
