package controllers

import (
	"net/http"
	"strconv"

	"demo/src/users/application"
	"github.com/gorilla/mux"
)

type DeleteUserController struct {
	DeleteUserUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUserUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{DeleteUserUseCase: deleteUserUseCase}
}

// Eliminar un usuario por ID
func (uc *DeleteUserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = uc.DeleteUserUseCase.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Usuario eliminado exitosamente"}`))
}
