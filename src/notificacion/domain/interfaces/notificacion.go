package interfaces

import "demo/src/notificacion/domain/entities"

type NotificacionRepository interface {
	CrearNotificacion(notificacion entities.Notificacion) error
	GetNotificaciones() ([]entities.Notificacion, error)
	GetNotificacionesResumidas() ([]entities.NotificacionResumida, error)  // Nuevo método
}
