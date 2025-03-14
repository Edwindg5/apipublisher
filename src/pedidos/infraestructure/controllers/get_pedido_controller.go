// api-database/src/pedidos/infraestructure/controllers/get_pedido_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"demo/src/pedidos/application"
	"github.com/gorilla/mux"
)



func ObtenerPedidosPendientes(useCase *application.GetPedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pedidos, err := useCase.ObtenerPedidosPendientes()
		if err != nil {
			http.Error(w, "Error al obtener los pedidos pendientes", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(pedidos)
	}
}

func BuscarPedidoPorID(useCase *application.GetPedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		pedido, err := useCase.BuscarPedidoPorID(id)
		if err != nil {
			http.Error(w, "Pedido no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(pedido)
	}
}
func ObtenerProductos(useCase *application.GetPedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productos, err := useCase.ObtenerTodosLosProductos()
		if err != nil {
			http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)
	}
}

func ObtenerPedidosPorCorreo(useCase *application.GetPedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		correo := r.URL.Query().Get("correo")
		if correo == "" {
			http.Error(w, "Correo es requerido", http.StatusBadRequest)
			return
		}

		pedidos, err := useCase.ObtenerPedidosPorCorreo(correo)
		if err != nil {
			http.Error(w, "Error al obtener pedidos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pedidos)
	}
}
