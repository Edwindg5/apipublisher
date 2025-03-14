package rabbitmq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
	"demo/src/pedidos/domain/entities"
)

func PublicarPedido(pedido entities.Pedido) error {
    rabbitURL := os.Getenv("RABBITMQ_URL")
    if rabbitURL == "" {
        log.Println("❌ ERROR: La variable de entorno RABBITMQ_URL no está configurada")
        return nil
    }

    log.Println("🔄 Conectando a RabbitMQ en:", rabbitURL)

    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        log.Println("❌ Error al conectar con RabbitMQ:", err)
        return err
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Println("❌ Error al abrir un canal en RabbitMQ:", err)
        return err
    }
    defer ch.Close()

    queue, err := ch.QueueDeclare(
        "pedidos_queue",
        true,  // Durable
        false, // AutoDelete
        false, // Exclusive
        false, // NoWait
        nil,   // Args
    )
    if err != nil {
        log.Println("❌ Error al declarar la cola:", err)
        return err
    }

    body, err := json.Marshal(pedido)
    if err != nil {
        log.Println("❌ Error al serializar el pedido:", err)
        return err
    }

    log.Println("📦 Enviando pedido a la cola:", queue.Name)

    err = ch.Publish(
        "",
        queue.Name,
        false,
        false,
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         body,
            DeliveryMode: amqp.Persistent,
        },
    )
    if err != nil {
        log.Println("❌ Error al publicar el pedido en RabbitMQ:", err)
        return err
    }

    log.Println("✅ Pedido enviado a RabbitMQ correctamente")
    return nil
}
