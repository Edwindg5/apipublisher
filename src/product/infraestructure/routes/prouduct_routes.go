package routes

import (
    "demo/src/product/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", controller.CreateProduct)
    }
}
