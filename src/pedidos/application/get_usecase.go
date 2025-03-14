// api-database/src/pedidos/application/get_usecase.go
package application

import (
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/repositories"
)

type GetPedidoUseCase struct {
	Repo repositories.GetPedidoRepository
}

func (uc *GetPedidoUseCase) ObtenerPedidosPendientes() ([]entities.Pedido, error) {
	return uc.Repo.ObtenerPedidosPendientes()
}

func (uc *GetPedidoUseCase) BuscarPedidoPorID(id int) (entities.Pedido, error) {
	return uc.Repo.BuscarPedidoPorID(id)
}
func (uc *GetPedidoUseCase) ObtenerTodosLosProductos() ([]entities.Pedido, error) {
	return uc.Repo.ObtenerTodosLosProductos()
}

func (uc *GetPedidoUseCase) ObtenerPedidosPorCorreo(correo string) ([]entities.Pedido, error) {
	return uc.Repo.ObtenerPedidosPorCorreo(correo)
}

