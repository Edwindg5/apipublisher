package application

import "demo/src/product/domain/interface"

type DeleteProductUsecase struct {
    Repo interfaces.ProductRepository
}

func NewDeleteProductUsecase(repo interfaces.ProductRepository) *DeleteProductUsecase {
    return &DeleteProductUsecase{Repo: repo}
}

func (uc *DeleteProductUsecase) DeleteProduct(id int) error {
    return uc.Repo.Delete(id)
}
