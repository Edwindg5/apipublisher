package application

import (
	"demo/src/product/domain/entities"
	"demo/src/product/infraestructure/repositories"
	"fmt"
)

type UpdateProductUsecase struct {
    Repo repositories.ProductRepository
}

func NewUpdateProductUsecase(repo repositories.ProductRepository) *UpdateProductUsecase {
    return &UpdateProductUsecase{Repo: repo}
}

func (uc *UpdateProductUsecase) UpdateProduct(product *entities.Product) error {
    existingProduct, err := uc.Repo.GetByID(product.ID)
    if err != nil {
        return err
    }
    if existingProduct == nil {
        return fmt.Errorf("producto con ID %d no encontrado", product.ID)
    }
    return uc.Repo.Update(product)
}
