package repositories

import (
	"database/sql"
	"demo/src/pedidos/domain/entities"
)

type CreatePedidoRepository struct {
	DB *sql.DB
}

func NewCreatePedidoRepository(db *sql.DB) *CreatePedidoRepository {
	return &CreatePedidoRepository{DB: db}
}

func (repo *CreatePedidoRepository) GuardarPedido(pedido entities.Pedido) error {
	_, err := repo.DB.Exec("INSERT INTO pedidos (cliente, producto, cantidad, estado) VALUES (?, ?, ?, ?)",
    pedido.Cliente, pedido.Producto, pedido.Cantidad, pedido.Estado)

	return err
}
