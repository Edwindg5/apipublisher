package controllers

import (
	"encoding/json"
	"net/http"
	"demo/src/users/application"
	"demo/src/users/domain/entities"
)

type UserController struct {
	UseCase *application.RegisterUserUseCase
}

func NewUserController(useCase *application.RegisterUserUseCase) *UserController {
	return &UserController{UseCase: useCase}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Entrada inválida", http.StatusBadRequest)
		return
	}

	if err := uc.UseCase.Execute(user); err != nil {
		http.Error(w, "Error al registrar usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado exitosamente"})
}
