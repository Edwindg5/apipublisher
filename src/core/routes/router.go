// api-database/src/core/routes/router.go
package routes

import (
	"database/sql"
	"demo/src/core"
	"demo/src/pedidos/infraestructure/routes"
	"net/http"

	"github.com/gorilla/mux"
)
func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(core.CORSMiddleware)    
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🚀 API funcionando correctamente"))
	}).Methods("GET")

	routes.RegisterPedidoRoutes(router, db)
	return router
}
