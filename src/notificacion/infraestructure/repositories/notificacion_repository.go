// api-database/src/notificacion/infraestructure/repositories/notificacion_repository.go
package repositories

import (
	"database/sql"
	"demo/src/notificacion/domain/entities"
)

type NotificacionRepository struct {
	DB *sql.DB
}

func NewNotificacionRepository(db *sql.DB) *NotificacionRepository {
	return &NotificacionRepository{DB: db}
}

func (repo *NotificacionRepository) GuardarNotificacion(notificacion entities.Notificacion) error {
	_, err := repo.DB.Exec(
		"INSERT INTO notificaciones (pedido_id, cliente, producto, cantidad, estado, fecha) VALUES (?, ?, ?, ?, ?, ?)",
		notificacion.PedidoID, notificacion.Cliente, notificacion.Producto, notificacion.Cantidad, notificacion.Estado, notificacion.Fecha,
	)
	return err
}
