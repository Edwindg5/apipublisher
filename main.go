package main

import (
	"demo/src/core"
	"demo/src/core/routes"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, verificando variables del sistema")
	}

	// Verificar la URL de RabbitMQ
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("❌ ERROR: La variable de entorno RABBITMQ_URL no está configurada")
	}
	log.Println("🔍 RabbitMQ URL cargada correctamente")

	// Conectar a la base de datos
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("❌ Error al conectar con la base de datos:", err)
	}
	defer db.Close()

	// Inicializar el router
	router := routes.SetupRouter(db)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}
