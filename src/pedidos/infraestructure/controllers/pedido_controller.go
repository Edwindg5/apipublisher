// api-database/src/pedidos/infraestructure/controllers/pedido_controller.go
package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"demo/src/pedidos/application"
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/rabbitmq"

	"github.com/gorilla/mux"
)

// CrearPedido almacena un nuevo pedido en la base de datos
func CrearPedido(pedidoUseCase *application.PedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pedido entities.Pedido
		if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
			http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
			return
		}

		// Guardar pedido en DB y enviarlo a RabbitMQ
		if err := pedidoUseCase.CrearPedido(pedido); err != nil {
			http.Error(w, "Error al guardar el pedido", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pedido)
	}
}

// BuscarPedidoPorNombre busca un pedido y lo envía a RabbitMQ
func BuscarPedidoPorNombre(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nombre := vars["nombre"]

		var pedido entities.Pedido
		err := db.QueryRow("SELECT id, cliente, producto, cantidad, estado FROM pedidos WHERE producto = ?", nombre).
			Scan(&pedido.ID, &pedido.Cliente, &pedido.Producto, &pedido.Cantidad, &pedido.Estado)

		if err == sql.ErrNoRows {
			http.Error(w, "Pedido no encontrado", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Error al buscar pedido", http.StatusInternalServerError)
			return
		}

		// Publicar en RabbitMQ
		if err := rabbitmq.PublicarPedido(pedido); err != nil {
			http.Error(w, "Error al enviar mensaje a RabbitMQ", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(pedido)
	}
}
