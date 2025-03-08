// api-database/src/pedidos/application/pedido_usecase.go
package application

import (
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/domain/interfaces"
	"demo/src/pedidos/infraestructure/rabbitmq"
)

type PedidoUseCase struct {
	Repo interfaces.PedidoRepository
}

func (p *PedidoUseCase) CrearPedido(pedido entities.Pedido) error {
	// Guardar el pedido en la base de datos
	err := p.Repo.GuardarPedido(pedido)
	if err != nil {
		return err
	}

	// Publicar el pedido en RabbitMQ
	return rabbitmq.PublicarPedido(pedido)
}

func (p *PedidoUseCase) ListarPedidos() ([]entities.Pedido, error) {
	return p.Repo.ObtenerPedidos()
}