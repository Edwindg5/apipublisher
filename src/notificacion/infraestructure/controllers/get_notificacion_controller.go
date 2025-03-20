// src/notificacion/infraestructure/controllers/get_notificacion_controller.go
package controllers

import (
    "encoding/json"
    "net/http"
    "demo/src/notificacion/application"
)

func ObtenerNotificaciones(useCase *application.GetNotificacionesUseCase) http.HandlerFunc {
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
