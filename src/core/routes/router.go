package routes

import (
	"database/sql"
	"net/http"

	userRoutes "demo/src/users/infraestructure/routes"
	productRoutes "demo/src/product/infraestructure/routes"
	"demo/src/product/application"
	"demo/src/product/infraestructure/controllers"
	"demo/src/product/infraestructure/repositories"

	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) http.Handler {
	// Crear router principal con `mux`
	mainRouter := mux.NewRouter()

	// Registrar rutas de usuarios con `mux`
	userRoutes.RegisterUserRoutes(mainRouter, db)

	// Crear router de `gin` para productos
	ginRouter := gin.Default() // Se usa `Default()` para logs y recovery
	productRepo := repositories.NewProductRepository(db)
	productUsecase := application.NewManageProductsUsecase(productRepo)
	productController := controllers.NewProductController(productUsecase)

	// Registrar rutas de productos en `gin`
	productRoutes.RegisterProductRoutes(ginRouter, productController)

	// Adaptar `ginRouter` para `mux`
	mainRouter.PathPrefix("/products").Handler(http.StripPrefix("/products", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ginRouter.ServeHTTP(w, r)
	})))

	return mainRouter
}
