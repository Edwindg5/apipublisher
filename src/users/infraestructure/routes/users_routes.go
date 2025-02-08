package routes

import (
	"database/sql"
	"demo/src/users/application"
	"demo/src/users/infraestructure/controllers"
	"demo/src/users/infraestructure/repositories"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes configura las rutas para usuarios
func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	repo := repositories.NewUserRepository(db)

	registerUseCase := application.NewRegisterUserUseCase(repo)
	getUserUseCase := application.NewGetUserUseCase(repo)
	putUserUseCase := application.NewPutUserUseCase(repo)
	deleteUserUseCase := application.NewDeleteUserUseCase(repo)

	registerController := controllers.NewUserController(registerUseCase)
	getUserController := controllers.NewGetUserController(getUserUseCase)
	putUserController := controllers.NewPutUserController(putUserUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

	userRoutes := router.PathPrefix("/api/v1/users").Subrouter()

	// Registrar usuario (POST)
	userRoutes.HandleFunc("", registerController.RegisterUser).Methods("POST", "OPTIONS")

	// Obtener usuarios (GET)
	userRoutes.HandleFunc("", getUserController.GetAllUsers).Methods("GET")
	userRoutes.HandleFunc("/{id}", getUserController.GetUserByID).Methods("GET")

	// Actualizar usuario (PUT)
	userRoutes.HandleFunc("/{id}", putUserController.UpdateUser).Methods("PUT").Headers("Content-Type", "application/json")

	// Eliminar usuario (DELETE)
	userRoutes.HandleFunc("/{id}", deleteUserController.DeleteUser).Methods("DELETE")
}
