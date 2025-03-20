// api-database/src/notificacion/infraestructure/repositories/notificacion_repository.go
package repositories

import (
	"database/sql"
	"demo/src/notificacion/domain/entities"
	"demo/src/notificacion/domain/interfaces" // Agregar la importación
)

type NotificacionRepository struct {
	DB *sql.DB
}

// 🔹 Añadir esta línea para asegurar que implementa la interfaz
var _ interfaces.NotificacionRepository = (*NotificacionRepository)(nil)

func NewNotificacionRepository(db *sql.DB) *NotificacionRepository {
	return &NotificacionRepository{DB: db}
}

func (repo *NotificacionRepository) CrearNotificacion(notificacion entities.Notificacion) error {
	_, err := repo.DB.Exec(
		"INSERT INTO notificaciones (pedido_id, cliente, producto, cantidad, estado, fecha) VALUES (?, ?, ?, ?, ?, ?)",
		notificacion.PedidoID, notificacion.Cliente, notificacion.Producto, notificacion.Cantidad, notificacion.Estado, notificacion.Fecha,
	)
	return err
}

func (repo *NotificacionRepository) GetNotificaciones() ([]entities.Notificacion, error) {
	rows, err := repo.DB.Query("SELECT id, pedido_id, cliente, producto, cantidad, estado, fecha FROM notificaciones")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var n entities.Notificacion
		if err := rows.Scan(&n.ID, &n.PedidoID, &n.Cliente, &n.Producto, &n.Cantidad, &n.Estado, &n.Fecha); err != nil {
			return nil, err
		}
		notificaciones = append(notificaciones, n)
	}
	return notificaciones, nil
}
