package controllers

import (
    "demo/src/product/application"
    "demo/src/product/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ProductController struct {
    Usecase *application.ManageProductsUsecase
}

func NewProductController(usecase *application.ManageProductsUsecase) *ProductController {
    return &ProductController{Usecase: usecase}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.Usecase.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el producto"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Producto registrado con éxito"})
}
