package repositories

import (
	"database/sql"
	"demo/src/pedidos/domain/entities"
)

type GetPedidoRepository struct {
	DB *sql.DB
}

func NewGetPedidoRepository(db *sql.DB) *GetPedidoRepository {
	return &GetPedidoRepository{DB: db}
}

func (repo *GetPedidoRepository) ObtenerPedidosPendientes() ([]entities.Pedido, error) {
	rows, err := repo.DB.Query("SELECT id, cliente, producto, cantidad, estado FROM pedidos WHERE estado = 'pendiente'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []entities.Pedido
	for rows.Next() {
		var p entities.Pedido
		err := rows.Scan(&p.ID, &p.Cliente, &p.Producto, &p.Cantidad, &p.Estado)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}


func (repo *GetPedidoRepository) BuscarPedidoPorID(id int) (entities.Pedido, error) {
	var pedido entities.Pedido
	err := repo.DB.QueryRow("SELECT id, cliente, producto, cantidad, estado FROM pedidos WHERE id = ?", id).
		Scan(&pedido.ID, &pedido.Cliente, &pedido.Producto, &pedido.Cantidad, &pedido.Estado)

	if err == sql.ErrNoRows {
		return pedido, nil
	}
	return pedido, err
}
func (r *GetPedidoRepository) ObtenerTodosLosProductos() ([]entities.Pedido, error) {
	rows, err := r.DB.Query("SELECT id, cliente, producto, cantidad, estado FROM pedidos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Pedido
	for rows.Next() {
		var pedido entities.Pedido
		err := rows.Scan(&pedido.ID, &pedido.Cliente, &pedido.Producto, &pedido.Cantidad, &pedido.Estado)
		if err != nil {
			return nil, err
		}
		productos = append(productos, pedido)
	}
	return productos, nil
}
