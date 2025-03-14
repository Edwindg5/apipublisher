// api-database/src/pedidos/application/create_usecase.go
package application

import (
	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/repositories"
)

type CreatePedidoUseCase struct {
	Repo repositories.CreatePedidoRepository
}   

func (uc *CreatePedidoUseCase) CrearPedido(pedido entities.Pedido) error {
	return uc.Repo.GuardarPedido(pedido)
}
