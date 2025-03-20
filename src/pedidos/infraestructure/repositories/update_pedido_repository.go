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
    // Consultar la cantidad real antes de actualizar
    var cantidadReal int
    querySelect := "SELECT cantidad FROM pedidos WHERE producto = ?"
    err := r.DB.QueryRow(querySelect, pedido.Producto).Scan(&cantidadReal)
    if err != nil {
        return fmt.Errorf("No se encontró un pedido con Producto: %s", pedido.Producto)
    }

    // Verificar si la cantidad enviada coincide con la cantidad real
    if pedido.Cantidad != cantidadReal {
        return fmt.Errorf("Cantidad incorrecta. Real: %d, Enviada: %d", cantidadReal, pedido.Cantidad)
    }

    // Usar la cantidad correcta en la actualización
    queryUpdate := "UPDATE pedidos SET estado = ? WHERE producto = ? AND cantidad = ?"
    result, err := r.DB.Exec(queryUpdate, pedido.Estado, pedido.Producto, pedido.Cantidad)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("No se encontró un pedido con Producto: %s y Cantidad: %d", pedido.Producto, pedido.Cantidad)
    }

    return nil
}



