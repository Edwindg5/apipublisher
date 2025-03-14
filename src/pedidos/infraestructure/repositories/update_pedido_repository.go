// api-database/src/pedidos/infraestructure/repositories/update_pedido_repository.go
package repositories

import (
	"database/sql"
	"demo/src/pedidos/domain/entities"
	"fmt"
)

type UpdatePedidoRepository struct {
	DB *sql.DB
}

func NewUpdatePedidoRepository(db *sql.DB) *UpdatePedidoRepository {
	return &UpdatePedidoRepository{DB: db}
}

func (r *UpdatePedidoRepository) UpdatePedido(pedido entities.Pedido) error {
    query := "UPDATE pedidos SET estado = ? WHERE producto = ? AND cantidad = ?"

    result, err := r.DB.Exec(query, pedido.Estado, pedido.Producto, pedido.Cantidad)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("No se encontró un pedido con Producto: %s y Cantidad: %d", pedido.Producto, pedido.Cantidad)
    }

    return nil
}

