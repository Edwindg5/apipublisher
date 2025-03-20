package application

import (
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/repositories"
	"demo/src/pedidos/infraestructure/rabbitmq" // Importar RabbitMQ
	"log"
)

type UpdatePedidoUseCase struct {
	Repo repositories.UpdatePedidoRepository
}

func (uc *UpdatePedidoUseCase) Execute(pedido entities.Pedido) error {
	// Actualizar el pedido en la base de datos
	err := uc.Repo.UpdatePedido(pedido)
	if err != nil {
		return err
	}

	// 🔥 Enviar a RabbitMQ después de actualizar
	log.Println("📦 Enviando pedido actualizado a RabbitMQ...")

	err = rabbitmq.PublicarPedido(pedido)
	if err != nil {
		log.Println("❌ Error al publicar el pedido en RabbitMQ:", err)
		return err
	}

	log.Println("✅ Pedido enviado a RabbitMQ correctamente")
	return nil
}
