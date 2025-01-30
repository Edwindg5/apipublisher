package controllers

import (
    "demo/src/product/application"
    "demo/src/product/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
   
)

type ProductController struct {
    CreateUsecase *application.CreateProductsUsecase
   
}

func NewProductController(createUC *application.CreateProductsUsecase,  ) *ProductController {
    return &ProductController{
        CreateUsecase: createUC,
    
    }
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := pc.CreateUsecase.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el producto"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Producto creado"})
}



