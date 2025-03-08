// api-database/src/pedidos/infraestructure/repositories/pedido_repository.go
package repositories

import (
	"database/sql"
	"demo/src/pedidos/domain/entities"
)

type PedidoRepository struct {
	DB *sql.DB
}

func NewPedidoRepository(db *sql.DB) *PedidoRepository {
	return &PedidoRepository{DB: db}
}

func (repo *PedidoRepository) GuardarPedido(pedido entities.Pedido) error {
	_, err := repo.DB.Exec("INSERT INTO pedidos (cliente, producto, cantidad, estado) VALUES (?, ?, ?, 'pendiente')",
		pedido.Cliente, pedido.Producto, pedido.Cantidad)
	return err
}

func (repo *PedidoRepository) ObtenerPedidos() ([]entities.Pedido, error) {
	rows, err := repo.DB.Query("SELECT id, cliente, producto, cantidad, estado FROM pedidos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []entities.Pedido
	for rows.Next() {
		var p entities.Pedido
		rows.Scan(&p.ID, &p.Cliente, &p.Producto, &p.Cantidad, &p.Estado)
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}

func (repo *PedidoRepository) BuscarPedidoPorNombre(nombre string) (entities.Pedido, error) {
	var pedido entities.Pedido
	err := repo.DB.QueryRow("SELECT id, cliente, producto, cantidad, estado FROM pedidos WHERE producto = ?", nombre).
		Scan(&pedido.ID, &pedido.Cliente, &pedido.Producto, &pedido.Cantidad, &pedido.Estado)

	if err == sql.ErrNoRows {
		return pedido, nil
	}
	return pedido, err
}
