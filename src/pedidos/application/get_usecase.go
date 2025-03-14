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
