package routes

import (
    "demo/src/product/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, 
    createcontroller *controllers.ProductController, 
    getController *controllers.GetProductController,
    updateController *controllers.UpdateProductController,
    deleteController *controllers.DeleteProductController) {

    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", createcontroller.CreateProduct)
        productRoutes.GET("/:id", getController.GetProduct)
        productRoutes.GET("/", getController.GetAllProducts)

        productRoutes.PUT("/:id", updateController.UpdateProduct)
        productRoutes.DELETE("/:id", deleteController.DeleteProduct)
    }
}
