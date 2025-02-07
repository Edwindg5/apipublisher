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

	// APLICAR MIDDLEWARE DE CORS PARA TODAS LAS RUTAS
	mainRouter.Use(core.MuxCORSMiddleware)

	// RUTA GLOBAL PARA MANEJAR SOLICITUDES OPTIONS
	mainRouter.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	// Registrar rutas de usuarios
	userRoutes.RegisterUserRoutes(mainRouter, db)
	

	// Configuración de Gin para productos
	ginRouter := gin.Default()
	ginRouter.Use(core.GinCORSMiddleware()) // Middleware para Gin

	productRepo := repositories.NewProductRepository(db)
	getProductRepo := repositories.NewGetProductRepository(db)
	updateProductRepo := repositories.NewUpdateProductRepository(db)
	deleteProductRepo := repositories.NewDeleteProductRepository(db)

	productUsecase := application.NewCreateProductsUsecase(productRepo)
	getProductUsecase := application.NewGetProductUsecase(getProductRepo)
	updateProductUsecase := application.NewUpdateProductUsecase(updateProductRepo)
	deleteProductUsecase := application.NewDeleteProductUsecase(deleteProductRepo)

	productController := controllers.NewProductController(productUsecase)
	getProductController := controllers.NewGetProductController(getProductUsecase)
	updateProductController := controllers.NewUpdateProductController(updateProductUsecase)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUsecase)

	productRoutes.RegisterProductRoutes(ginRouter, productController, getProductController, updateProductController, deleteProductController)

	mainRouter.PathPrefix("/products").Handler(http.StripPrefix("/products", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ginRouter.ServeHTTP(w, r)
	})))

	return mainRouter
}
