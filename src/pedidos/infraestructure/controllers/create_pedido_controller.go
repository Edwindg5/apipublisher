package controllers

import (
	"encoding/json"
	"net/http"

	"demo/src/pedidos/application"
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/rabbitmq"
)

func CrearPedido(useCase *application.CreatePedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pedido entities.Pedido
		if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
			http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
			return
		}

		// 1️⃣ Guardar el pedido en la base de datos
		if err := useCase.CrearPedido(pedido); err != nil {
			http.Error(w, "Error al guardar el pedido", http.StatusInternalServerError)
			return
		}

		// 2️⃣ Publicar el pedido en RabbitMQ
		if err := rabbitmq.PublicarPedido(pedido); err != nil {
			http.Error(w, "Error al enviar el pedido a RabbitMQ", http.StatusInternalServerError)
			return
		}

		// 3️⃣ Responder con éxito
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Pedido enviado a la cola correctamente"})
	}
}
