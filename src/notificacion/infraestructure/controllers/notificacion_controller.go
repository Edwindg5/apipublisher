// src/notificacion/infraestructure/controllers/notificacion_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"demo/src/notificacion/application"
	"demo/src/notificacion/domain/entities"
)

func CrearNotificacion(useCase *application.CreateNotificacionUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notificacion entities.Notificacion

		if err := json.NewDecoder(r.Body).Decode(&notificacion); err != nil {
			http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
			return
		}

		if notificacion.PedidoID <= 0 || notificacion.Cliente == "" || notificacion.Producto == "" {
			http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
			return
		}

		if err := useCase.CrearNotificacion(notificacion); err != nil {
			http.Error(w, "Error al guardar la notificación", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Notificación guardada correctamente"})
	}
}
