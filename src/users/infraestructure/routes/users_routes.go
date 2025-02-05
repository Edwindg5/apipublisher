package routes

import (
	"database/sql"
	"demo/src/users/infraestructure/controllers"
	"demo/src/users/infraestructure/repositories"
	"demo/src/users/application"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	repo := repositories.NewUserRepository(db)
	useCase := application.NewRegisterUserUseCase(repo)
	controller := controllers.NewUserController(useCase)

	// ACEPTAR POST Y OPTIONS
	router.HandleFunc("/users/register", controller.RegisterUser).Methods("POST", "OPTIONS")
}
