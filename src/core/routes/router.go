package routes

import (
	"database/sql"
	"net/http"

	userRoutes "demo/src/users/infraestructure/routes"
	productRoutes "demo/src/product/infraestructure/routes"
	"demo/src/product/application"
	"demo/src/product/infraestructure/controllers"
	"demo/src/product/infraestructure/repositories"
	"demo/src/core"

	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)


func NewRouter(db *sql.DB) http.Handler {
	mainRouter := mux.NewRouter()

	
	mainRouter.Use(core.MuxCORSMiddleware)


	mainRouter.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})


	userRoutes.RegisterUserRoutes(mainRouter, db)

	
	ginRouter := gin.Default()
	ginRouter.Use(core.GinCORSMiddleware()) // Middleware para Gin

	productRepo := repositories.NewProductRepository(db)

	
	productUsecase := application.NewCreateProductsUsecase(productRepo)
	getProductUsecase := application.NewGetProductUsecase(productRepo)
	updateProductUsecase := application.NewUpdateProductUsecase(productRepo)
	deleteProductUsecase := application.NewDeleteProductUsecase(productRepo)


	productController := controllers.NewProductController(productUsecase)
	getProductController := controllers.NewGetProductController(getProductUsecase)
	updateProductController := controllers.NewUpdateProductController(updateProductUsecase)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUsecase)


	api := ginRouter.Group("/api/v1")
	productRoutes.RegisterProductRoutes(api, productController, getProductController, updateProductController, deleteProductController)


	mainRouter.PathPrefix("/api/v1").Handler(ginRouter)

	return mainRouter
}
