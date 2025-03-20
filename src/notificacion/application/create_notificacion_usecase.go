// api-database/src/notificacion/application/create_notificacion_usecase.go
package application

import (
	"demo/src/notificacion/domain/entities"
	"demo/src/notificacion/infraestructure/repositories"
)

type CreateNotificacionUseCase struct {
	Repo repositories.NotificacionRepository
}

func (uc *CreateNotificacionUseCase) CrearNotificacion(notificacion entities.Notificacion) error {
	return uc.Repo.GuardarNotificacion(notificacion)
}
