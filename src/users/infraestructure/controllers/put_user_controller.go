package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"github.com/gorilla/mux"
)

type PutUserController struct {
	PutUserUseCase *application.PutUserUseCase
}

func NewPutUserController(putUserUseCase *application.PutUserUseCase) *PutUserController {
	return &PutUserController{PutUserUseCase: putUserUseCase}
}

// Actualizar un usuario por ID
func (uc *PutUserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Configurar respuesta JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtener el ID desde la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	// Decodificar el cuerpo de la solicitud
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Entrada inválida"}`, http.StatusBadRequest)
		return
	}

	// Asignar ID al usuario
	user.ID = id
	err = uc.PutUserUseCase.UpdateUser(user)
	if err != nil {
		http.Error(w, `{"error": "Error al actualizar usuario"}`, http.StatusInternalServerError)
		return
	}

	// Enviar respuesta JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario actualizado correctamente"})
}
