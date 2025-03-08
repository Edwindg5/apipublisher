//api-database/src/pedidos/domain/interfaces/pedido_repository.go
package interfaces

import "demo/src/pedidos/domain/entities"

type PedidoRepository interface {
	GuardarPedido(pedido entities.Pedido) error
	ObtenerPedidos() ([]entities.Pedido, error)
}
