package routes

import (
	"demo/src/users/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registra las rutas para los usuarios
func RegisterUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
}
