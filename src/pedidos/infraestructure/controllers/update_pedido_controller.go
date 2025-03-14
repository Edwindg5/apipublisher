package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"demo/src/pedidos/application"
	"demo/src/pedidos/domain/entities"
)

func ActualizarPedido(useCase *application.UpdatePedidoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pedido entities.Pedido
		err := json.NewDecoder(r.Body).Decode(&pedido)
		if err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		log.Printf("📥 Pedido recibido para actualizar: %+v", pedido)

		pedido.Estado = "procesado" // Marcar como procesado

		err = useCase.Execute(pedido)
		if err != nil {
			log.Println("❌ Error al actualizar pedido:", err)
			http.Error(w, "Error al actualizar pedido", http.StatusInternalServerError)
			return
		}

		log.Printf("✅ Pedido actualizado: %+v", pedido)

		// 🔥 Notificar a los clientes SSE
		NotificarPedidoActualizado(pedido)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Pedido actualizado",
		})
	}
}
