// application/deleteProduct_usecase.go
package application

import "demo/src/product/infraestructure/repositories"

type DeleteProductUsecase struct {
    Repo repositories.ProductRepository
}

func NewDeleteProductUsecase(repo repositories.ProductRepository) *DeleteProductUsecase {
    return &DeleteProductUsecase{Repo: repo}
}

func (uc *DeleteProductUsecase) DeleteProduct(id int) error {
    return uc.Repo.Delete(id)
}