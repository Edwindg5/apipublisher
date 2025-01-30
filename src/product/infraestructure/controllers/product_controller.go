package controllers

import (
    "demo/src/product/application"
    "demo/src/product/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type ProductController struct {
    CreateUsecase *application.ManageProductsUsecase
    GetUsecase    *application.GetProductUsecase
    UpdateUsecase *application.UpdateProductUsecase
    DeleteUsecase *application.DeleteProductUsecase
}

func NewProductController(createUC *application.ManageProductsUsecase, getUC *application.GetProductUsecase, updateUC *application.UpdateProductUsecase, deleteUC *application.DeleteProductUsecase) *ProductController {
    return &ProductController{
        CreateUsecase: createUC,
        GetUsecase:    getUC,
        UpdateUsecase: updateUC,
        DeleteUsecase: deleteUC,
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


func (pc *ProductController) GetProduct(c *gin.Context) {
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

func (pc *ProductController) UpdateProduct(c *gin.Context) {
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
    
    product.ID = id // Asegurar que el ID en la URL se usa correctamente

    if err := pc.UpdateUsecase.UpdateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    if err := pc.DeleteUsecase.DeleteProduct(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
