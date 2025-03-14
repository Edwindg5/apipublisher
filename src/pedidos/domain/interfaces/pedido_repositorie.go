// api-database/src/pedidos/domain/interfaces/pedido_repository.go
package interfaces

import "demo/src/pedidos/domain/entities"

type PedidoRepository interface {
	GuardarPedido(pedido entities.Pedido) error
	ObtenerPedidosPendientes() ([]entities.Pedido, error)
	BuscarPedidoPorID(id int) (entities.Pedido, error)
	ActualizarPedidoPorProducto(pedido entities.Pedido) error
	ObtenerTodosLosProductos() ([]entities.Pedido, error)
	ObtenerPedidosPorCorreo(correo string) ([]entities.Pedido, error)

	ObtenerPedidosPendientesPorProducto(nombreProducto string) ([]entities.Pedido, error)
	CambiarEstadoPedido(id int, estado string) error

	ActualizarEstadoYReducirCantidad(id int, nuevoEstado string) error
	EliminarPedido(id int) error
}
