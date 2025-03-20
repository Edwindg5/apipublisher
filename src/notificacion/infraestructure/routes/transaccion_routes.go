//api-database/src/notificacion/infraestructure/routes/notificacion_routes.go
package routes

import (
	"database/sql"
	"demo/src/notificacion/application"
	"demo/src/notificacion/infraestructure/controllers"
	"demo/src/notificacion/infraestructure/repositories"

	"github.com/gorilla/mux"
)

func RegisterNotificacionRoutes(router *mux.Router, db *sql.DB) {
	notificacionRepo := repositories.NewNotificacionRepository(db)
	notificacionUseCase := &application.CreateNotificacionUseCase{Repo: *notificacionRepo}

	router.HandleFunc("/notificaciones", controllers.CrearNotificacion(notificacionUseCase)).Methods("POST")
}
