package routes

import (
	"database/sql"
	"demo/src/notificacion/application"
	"demo/src/notificacion/domain/interfaces"
	"demo/src/notificacion/infraestructure/controllers"
	"demo/src/notificacion/infraestructure/repositories"

	"github.com/gorilla/mux"
)

func RegisterNotificacionRoutes(router *mux.Router, db *sql.DB) {
	notificacionRepo := repositories.NewNotificacionRepository(db)

	// Asignar la interfaz a la implementación
	var repoInterface interfaces.NotificacionRepository = notificacionRepo

	// Casos de uso
	notificacionUseCase := &application.CreateNotificacionUseCase{Repo: repoInterface}
	notificacionGetUseCase := &application.GetNotificacionesUseCase{Repo: repoInterface}
	notificacionGetResumidaUseCase := &application.GetNotificacionesResumidasUseCase{Repo: repoInterface}

	// Rutas
	router.HandleFunc("/notificaciones", controllers.CrearNotificacion(notificacionUseCase)).Methods("POST")
	router.HandleFunc("/notificaciones", controllers.ObtenerNotificaciones(notificacionGetUseCase)).Methods("GET")

	// Nueva ruta para notificaciones resumidas
	router.HandleFunc("/notificaciones/resumidas", controllers.ObtenerNotificacionesResumidas(notificacionGetResumidaUseCase)).Methods("GET")
}
