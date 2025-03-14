package main

import (
	"demo/src/core"
	"demo/src/core/routes"
	"demo/src/pedidos/application"
	"demo/src/pedidos/infraestructure/controllers"
	"demo/src/pedidos/infraestructure/repositories"

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

	// 🔹 Instanciar el repositorio y el caso de uso
	pedidoRepo := repositories.NewUpdatePedidoRepository(db)
	useCase := &application.UpdatePedidoUseCase{Repo: *pedidoRepo}

	// Configurar el router
	router := routes.SetupRouter(db)
	handler := core.CORSMiddleware(router) // Middleware aplicado aquí

	// 🔥 Registrar rutas SSE y actualización de pedido
	http.HandleFunc("/stream-pedidos", controllers.PedidosSSE)
	http.HandleFunc("/actualizar-pedido", controllers.ActualizarPedido(useCase))

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}
