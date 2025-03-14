package routes

import (
	"database/sql"
	"demo/src/pedidos/application"
	"demo/src/pedidos/infraestructure/controllers"
	"demo/src/pedidos/infraestructure/repositories"

	"github.com/gorilla/mux"
)

func RegisterPedidoRoutes(router *mux.Router, db *sql.DB) {
	createRepo := repositories.NewCreatePedidoRepository(db)
	getRepo := repositories.NewGetPedidoRepository(db)
	updateRepo := repositories.NewUpdatePedidoRepository(db)

	createUseCase := &application.CreatePedidoUseCase{Repo: *createRepo}
	getUseCase := &application.GetPedidoUseCase{Repo: *getRepo}
	putUseCase := &application.UpdatePedidoUseCase{Repo: *updateRepo}

	router.HandleFunc("/pedidos", controllers.CrearPedido(createUseCase)).Methods("POST")
	router.HandleFunc("/pedidos/pendientes", controllers.ObtenerPedidosPendientes(getUseCase)).Methods("GET")
	router.HandleFunc("/pedidos/{id:[0-9]+}", controllers.BuscarPedidoPorID(getUseCase)).Methods("GET")
	router.HandleFunc("/pedidos/{id:[0-9]+}", controllers.ActualizarPedido(putUseCase)).Methods("PUT")
	router.HandleFunc("/productos", controllers.ObtenerProductos(getUseCase)).Methods("GET")
	router.HandleFunc("/pedidos/correo/{correo}", controllers.ObtenerPedidosPorCorreo(getUseCase)).Methods("GET")
	router.HandleFunc("/pedidos/actualizar", controllers.ActualizarPedido(putUseCase)).Methods("PUT")
	router.HandleFunc("/stream-pedidos", controllers.PedidosSSE).Methods("GET")


}
