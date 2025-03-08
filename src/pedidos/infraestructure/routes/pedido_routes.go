package routes

import (
	"database/sql"
	"demo/src/pedidos/application"
	"demo/src/pedidos/infraestructure/controllers"
	"demo/src/pedidos/infraestructure/repositories"

	"github.com/gorilla/mux"
)

func RegisterPedidoRoutes(router *mux.Router, db *sql.DB) {
    pedidoRepo := repositories.NewPedidoRepository(db)
    pedidoUseCase := &application.PedidoUseCase{Repo: pedidoRepo}

    router.HandleFunc("/pedidos", controllers.CrearPedido(pedidoUseCase)).Methods("POST")
    router.HandleFunc("/pedidos/{nombre}", controllers.BuscarPedidoPorNombre(db)).Methods("GET")
}
