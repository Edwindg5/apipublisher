package main

import (
	"demo/src/core"
	"demo/src/users/infraestructure"
	"demo/src/routes"
	"demo/src/users/application"
	"demo/src/users/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Configuración de la base de datos
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// Repositorios y casos de uso
	userRepo := infrastructure.NewUserRepositoryImpl(db)
	registerUserUseCase := application.NewRegisterUserUseCase(userRepo)
	loginUserUseCase := application.NewLoginUserUseCase(userRepo)

	// Controladores
	userController := controllers.NewUserController(registerUserUseCase, loginUserUseCase)

	// Configuración de Gin
	router := gin.Default()

	// Middleware global (opcional, ej. CORS)
	router.Use(core.CORSMiddleware())

	// Rutas
	api := router.Group("/api")
	routes.RegisterUserRoutes(api, userController)

	// Arrancar el servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor corriendo en el puerto %s", port)
	log.Fatal(router.Run(":" + port))
}
