// api-database/src/pedidos/infraestructure/rabbitmq/publisher.go
package rabbitmq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
	"demo/src/pedidos/domain/entities"
)

func PublicarPedido(pedido entities.Pedido) error {
	// Obtener la URL de RabbitMQ desde las variables de entorno
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Println("❌ ERROR: La variable de entorno RABBITMQ_URL no está configurada")
		return nil
	}

	// Conectar a RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Println("❌ Error al conectar con RabbitMQ:", err)
		return err
	}
	log.Println("✅ Conexión exitosa con RabbitMQ 🚀")
	defer conn.Close()

	// Crear canal
	ch, err := conn.Channel()
	if err != nil {
		log.Println("❌ Error al abrir un canal en RabbitMQ:", err)
		return err
	}
	defer ch.Close()

	// Convertir el pedido a JSON
	body, err := json.Marshal(pedido)
	if err != nil {
		log.Println("❌ Error al serializar el pedido:", err)
		return err
	}

	log.Println("📤 Enviando mensaje a la cola 'pedidos_queue':", string(body))

	// Publicar mensaje en RabbitMQ (sin declarar la cola porque ya existe en RabbitMQ)
	err = ch.Publish(
		"", // Intercambio vacío (usa la cola por defecto)
		"pedidos_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // Mensaje persistente porque la cola es durable
		},
	)
	if err != nil {
		log.Println("❌ Error al publicar el pedido en RabbitMQ:", err)
		return err
	}

	log.Println("✅ Mensaje publicado correctamente en RabbitMQ 🚀")
	return nil
}
