package application

import (
	"demo/src/notificacion/domain/entities"
	"demo/src/notificacion/domain/interfaces"
)

type GetNotificacionesResumidasUseCase struct {
    Repo interfaces.NotificacionRepository
}

func (uc *GetNotificacionesResumidasUseCase) Execute() ([]entities.NotificacionResumida, error) {
	return uc.Repo.GetNotificacionesResumidas()
}
