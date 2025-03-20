// api-database/src/notificacion/infraestructure/routes/notificacion_routes.go
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

	// Ahora puedes asignarlo directamente a la interfaz
	var repoInterface interfaces.NotificacionRepository = notificacionRepo

	notificacionUseCase := &application.CreateNotificacionUseCase{Repo: repoInterface}
	notificacionGetUseCase := &application.GetNotificacionesUseCase{Repo: repoInterface}

	router.HandleFunc("/notificaciones", controllers.CrearNotificacion(notificacionUseCase)).Methods("POST")
	router.HandleFunc("/notificaciones", controllers.ObtenerNotificaciones(notificacionGetUseCase)).Methods("GET")
}


