package controllers

import (
	"net/http"
	"demo/src/core"
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"demo/src/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	RegisterUseCase *application.RegisterUserUseCase
	LoginUseCase    *application.LoginUserUseCase
}

func NewUserController(registerUC *application.RegisterUserUseCase, loginUC *application.LoginUserUseCase) *UserController {
	return &UserController{RegisterUseCase: registerUC, LoginUseCase: loginUC}
}

func (uc *UserController) Register(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		core.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Hash the password before registration
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		core.JSONResponse(c, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}
	user.Password = hashedPassword

	if err := uc.RegisterUseCase.Execute(&user); err != nil {
		core.JSONResponse(c, http.StatusInternalServerError, "Failed to register user", nil)
		return
	}

	// Auto-generate token after successful registration
	token, err := uc.LoginUseCase.GenerateToken(user.ID, user.Email)
	if err != nil {
		core.JSONResponse(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	core.JSONResponse(c, http.StatusCreated, "User registered successfully", gin.H{"token": token})
}

func (uc *UserController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		core.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	token, err := uc.LoginUseCase.Execute(credentials.Email, credentials.Password)
	if err != nil {
		core.JSONResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	core.JSONResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
