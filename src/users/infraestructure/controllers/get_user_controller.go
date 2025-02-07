package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"demo/src/users/application"
	"github.com/gorilla/mux"
)

type GetUserController struct {
	GetUserUseCase *application.GetUserUseCase
}

func NewGetUserController(getUserUseCase *application.GetUserUseCase) *GetUserController {
	return &GetUserController{GetUserUseCase: getUserUseCase}
}

// Obtener todos los usuarios
func (uc *GetUserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.GetUserUseCase.GetAllUsers()
	if err != nil {
		http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Obtener un usuario por ID
func (uc *GetUserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := uc.GetUserUseCase.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
