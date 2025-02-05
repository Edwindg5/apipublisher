// src/product/application/getProduct_usecase.go
package application

import (
    "demo/src/product/domain/entities"
    "demo/src/product/infraestructure/repositories"
)

type GetProductUsecase struct {
    Repo repositories.GetProductRepository
}

func NewGetProductUsecase(repo repositories.GetProductRepository) *GetProductUsecase {
    return &GetProductUsecase{Repo: repo}
}

func (uc *GetProductUsecase) GetProductByID(id int) (*entities.Product, error) {
    return uc.Repo.GetByID(id)
}

func (uc *GetProductUsecase) GetAllProducts() ([]*entities.Product, error) {
    return uc.Repo.GetAll()
}
