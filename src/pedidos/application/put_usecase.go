// api-database/src/pedidos/application/update_usecase.go
package application

import (

	"demo/src/pedidos/domain/entities"
	"demo/src/pedidos/infraestructure/repositories"
)

type UpdatePedidoUseCase struct {
	Repo repositories.UpdatePedidoRepository
}

func (uc *UpdatePedidoUseCase) Execute(pedido entities.Pedido) error {
	return uc.Repo.UpdatePedido(pedido)
}
