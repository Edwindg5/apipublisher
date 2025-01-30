package routes

import (
    "demo/src/product/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", controller.CreateProduct)
        productRoutes.GET("/:id", controller.GetProduct)
        productRoutes.PUT("/:id", controller.UpdateProduct)
        productRoutes.DELETE("/:id", controller.DeleteProduct)
    }
}