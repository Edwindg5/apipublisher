package application

import (
    "demo/src/product/domain/entities"
    "demo/src/product/infraestructure/repositories"
)

type ManageProductsUsecase struct {
    Repo repositories.ProductRepository
}

func NewManageProductsUsecase(repo repositories.ProductRepository) *ManageProductsUsecase {
    return &ManageProductsUsecase{Repo: repo}
}

func (uc *ManageProductsUsecase) CreateProduct(product *entities.Product) error {
    return uc.Repo.Create(product)
}
