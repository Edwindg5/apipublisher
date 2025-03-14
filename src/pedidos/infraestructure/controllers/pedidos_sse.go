package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"demo/src/pedidos/domain/entities"
)

var clients = make(map[chan string]bool) // Clientes SSE
var mutex = sync.Mutex{}

func PedidosSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	messageChan := make(chan string)
	mutex.Lock()
	clients[messageChan] = true
	mutex.Unlock()

	log.Println("👥 Cliente SSE conectado")

	for {
		select {
		case msg := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			mutex.Lock()
			delete(clients, messageChan)
			mutex.Unlock()
			log.Println("❌ Cliente SSE desconectado")
			return
		}
	}
}

func NotificarPedidoActualizado(pedido entities.Pedido) {
	mutex.Lock()
	defer mutex.Unlock()

	jsonData, err := json.Marshal(pedido)
	if err != nil {
		log.Println("❌ Error serializando pedido SSE:", err)
		return
	}

	for clientChan := range clients {
		log.Println("📡 Enviando actualización SSE:", string(jsonData))
		clientChan <- string(jsonData)
	}
}
