// api-database/src/notificacion/application/create_notificacion_usecase.go
package application

import (
	"demo/src/notificacion/domain/entities"
	"demo/src/notificacion/domain/interfaces" // Usar interfaces en vez de repositories
)

type CreateNotificacionUseCase struct {
	Repo interfaces.NotificacionRepository // Cambiar de repositories.NotificacionRepository a interfaces.NotificacionRepository
}

func (uc *CreateNotificacionUseCase) CrearNotificacion(notificacion entities.Notificacion) error {
	return uc.Repo.CrearNotificacion(notificacion)
}
