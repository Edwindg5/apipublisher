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

        log.Printf("📥 Recibida solicitud de actualización. Producto: %s, Cantidad: %d", pedido.Producto, pedido.Cantidad)

        pedido.Estado = "procesado"

        err = useCase.Execute(pedido)
        if err != nil {
            log.Println("❌ Error al actualizar el pedido en la base de datos:", err)
            http.Error(w, "Error al actualizar el pedido", http.StatusInternalServerError)
            return
        }

        log.Printf("✅ Pedido actualizado en la base de datos. Producto: %s, Cantidad: %d, Estado: %s", pedido.Producto, pedido.Cantidad, pedido.Estado)
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Pedido actualizado a 'procesado'"})
    }
}

