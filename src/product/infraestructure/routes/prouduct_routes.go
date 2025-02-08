package routes

import (
    "demo/src/product/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)


func RegisterProductRoutes(router *gin.RouterGroup, 
    productController *controllers.ProductController, 
    getProductController *controllers.GetProductController,
    updateProductController *controllers.UpdateProductController,
    deleteProductController *controllers.DeleteProductController) {

    productRoutes := router.Group("/products")
    {
        productRoutes.POST("", productController.CreateProduct) 
        productRoutes.GET("", getProductController.GetAllProducts) 
        productRoutes.GET("/:id", getProductController.GetProduct) 

        productRoutes.PUT("/:id", updateProductController.UpdateProduct) 
        productRoutes.DELETE("/:id", deleteProductController.DeleteProduct) 
    }
}
