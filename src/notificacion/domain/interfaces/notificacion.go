// api-database/src/notificacion/domain/interfaces/notificacion.go
package interfaces

import "demo/src/notificacion/domain/entities"

// En interfaces/notificacion.go
type NotificacionRepository interface {
	CrearNotificacion(notificacion entities.Notificacion) error
	GetNotificaciones() ([]entities.Notificacion, error)
}
