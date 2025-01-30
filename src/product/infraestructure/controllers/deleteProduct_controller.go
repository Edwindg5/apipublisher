// src/product/infraestructure/controllers/deleteProduct_controller.go
package controllers

import (
    "demo/src/product/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type DeleteProductController struct {
    DeleteUsecase *application.DeleteProductUsecase
}

func NewDeleteProductController(deleteUC *application.DeleteProductUsecase) *DeleteProductController {
    return &DeleteProductController{DeleteUsecase: deleteUC}
}

func (pc *DeleteProductController) DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := pc.DeleteUsecase.DeleteProduct(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
