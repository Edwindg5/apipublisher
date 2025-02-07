package routes

import (
	"database/sql"
	"demo/src/users/application"
	"demo/src/users/infraestructure/controllers"
	"demo/src/users/infraestructure/repositories"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	repo := repositories.NewUserRepository(db)
	getRepo := repositories.NewGetUserRepository(db)
	putRepo := repositories.NewPutUserRepository(db)
	deleteRepo := repositories.NewDeleteUserRepository(db)

	registerUseCase := application.NewRegisterUserUseCase(repo)
	getUserUseCase := application.NewGetUserUseCase(getRepo)
	putUserUseCase := application.NewPutUserUseCase(putRepo)
	deleteUserUseCase := application.NewDeleteUserUseCase(deleteRepo)

	registerController := controllers.NewUserController(registerUseCase)
	getUserController := controllers.NewGetUserController(getUserUseCase)
	putUserController := controllers.NewPutUserController(putUserUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

	// Registrar usuario (POST)
	router.HandleFunc("/users/register", registerController.RegisterUser).Methods("POST", "OPTIONS")

	// Obtener usuarios (GET)
	router.HandleFunc("/users", getUserController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUserController.GetUserByID).Methods("GET")

	// Actualizar usuario (PUT) - Especificar que es un JSON en el Header
	router.HandleFunc("/users/{id}", putUserController.UpdateUser).Methods("PUT").Headers("Content-Type", "application/json")

	// Eliminar usuario (DELETE)
	router.HandleFunc("/users/{id}", deleteUserController.DeleteUser).Methods("DELETE")
}
