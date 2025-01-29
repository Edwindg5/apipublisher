package main

import (
    "log"
    "net/http"
    "demo/src/core"
    "demo/src/core/routes"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error al cargar el archivo .env")
    }

    db, err := core.ConnectDB()
    if err != nil {
        log.Fatal("Error al conectar la base de datos:", err)
    }
    defer db.Close()

    router := routes.NewRouter(db)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Servidor corriendo en el puerto", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
