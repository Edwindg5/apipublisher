//src/product/infraestructure/controllers/updateProduct_controller.go
package controllers

import (
    "demo/src/product/application"
    "demo/src/product/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type UpdateProductController struct {
    UpdateUsecase *application.UpdateProductUsecase
}

func NewUpdateProductController(updateUC *application.UpdateProductUsecase) *UpdateProductController {
    return &UpdateProductController{UpdateUsecase: updateUC}
}

func (pc *UpdateProductController) UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product.ID = id

    if err := pc.UpdateUsecase.UpdateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}
