package application

import (
    "demo/src/product/domain/entities"
    "demo/src/product/domain/interface"
)

type GetProductUsecase struct {
    Repo interfaces.ProductRepository
}

func NewGetProductUsecase(repo interfaces.ProductRepository) *GetProductUsecase {
    return &GetProductUsecase{Repo: repo}
}

func (uc *GetProductUsecase) GetProductByID(id int) (*entities.Product, error) {
    return uc.Repo.GetByID(id)
}

func (uc *GetProductUsecase) GetAllProducts() ([]*entities.Product, error) {
    return uc.Repo.GetAll()
}
