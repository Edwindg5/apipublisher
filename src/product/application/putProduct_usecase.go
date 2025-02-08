package application

import (
	"demo/src/product/domain/entities"
	"demo/src/product/domain/interface"
	"fmt"
)

type UpdateProductUsecase struct {
    Repo interfaces.ProductRepository
}

func NewUpdateProductUsecase(repo interfaces.ProductRepository) *UpdateProductUsecase {
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
