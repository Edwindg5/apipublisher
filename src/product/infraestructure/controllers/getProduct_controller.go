// src/product/infraestructure/controllers/getProduct_controller.go
package controllers

import (
    "demo/src/product/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GetProductController struct {
    GetUsecase *application.GetProductUsecase
}

func NewGetProductController(getUC *application.GetProductUsecase) *GetProductController {
    return &GetProductController{GetUsecase: getUC}
}

func (pc *GetProductController) GetProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    product, err := pc.GetUsecase.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }
    c.JSON(http.StatusOK, product)
}