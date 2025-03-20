package controllers

import (
	"encoding/json"
	"net/http"

	"demo/src/pedidos/application"
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/rabbitmq"
	"strconv"
)

func CrearPedido(useCase *application.CreatePedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pedido entities.Pedido

		// Verifica si el JSON es válido
		if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
			http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
			return
		}

		// Verifica campos obligatorios
		if pedido.Cliente == "" || pedido.Producto == "" || pedido.Cantidad <= 0 || pedido.Estado == "" {
			http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
			return
		}

		// Guarda el pedido en la base de datos
		if err := useCase.CrearPedido(pedido); err != nil {
			http.Error(w, "Error al guardar el pedido", http.StatusInternalServerError)
			return
		}

		// Verifica si se debe enviar a la cola (flag opcional en la URL)
		enviarColaParam := r.URL.Query().Get("enviar_a_cola")
		enviarCola, _ := strconv.ParseBool(enviarColaParam)

		if enviarCola {
			// Publica el pedido en RabbitMQ si el flag es true
			if err := rabbitmq.PublicarPedido(pedido); err != nil {
				http.Error(w, "Error al enviar el pedido a RabbitMQ", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Pedido enviado a la cola correctamente"})
		} else {
			// No enviar a la cola, solo guardar en la BD
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Pedido guardado en la base de datos"})
		}
	}
}
