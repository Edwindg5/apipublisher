package routes

import (
	"database/sql"
	"net/http"
	"demo/src/pedidos/infraestructure/routes" 

	"github.com/gorilla/mux"
)

// SetupRouter configura las rutas y usa la base de datos
func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// Ruta de prueba
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🚀 API funcionando correctamente"))
	}).Methods("GET")

	// Registrar las rutas de pedidos
	routes.RegisterPedidoRoutes(router, db) // ✅ Llamada corregida

	return router
}

// Middleware de Logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("📝", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
