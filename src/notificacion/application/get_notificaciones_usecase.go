// api-database/src/notificacion/application/get_notificaciones_usecase.go
package application

import (
	"demo/src/notificacion/domain/entities"
	"demo/src/notificacion/domain/interfaces"
)

type GetNotificacionesUseCase struct {
    Repo interfaces.NotificacionRepository
}

func (uc *GetNotificacionesUseCase) Execute() ([]entities.Notificacion, error) {
	return uc.Repo.GetNotificaciones()
}
