package controllers

import (
	"encoding/json"
	"net/http"
	"demo/src/notificacion/application"
)

func ObtenerNotificacionesResumidas(useCase *application.GetNotificacionesResumidasUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notificaciones, err := useCase.Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notificaciones)
	}
}
