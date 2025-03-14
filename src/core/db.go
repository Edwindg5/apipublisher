// api-database/src/core/db.go
package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, verificando variables del sistema")
	}
}


func ConnectDB() (*sql.DB, error) {
	loadEnv()

	log.Println("📌 Cargando configuración de la base de datos:")
	log.Println("DB_USER:", os.Getenv("DB_USER"))
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", getEnv("DB_PORT", "3306"))
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		getEnv("DB_PORT", "3306"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("❌ Error al abrir la conexión a MySQL:", err)
		return nil, err
	}


	if err := db.Ping(); err != nil {
		log.Println("❌ Error al hacer ping a MySQL:", err)
		return nil, err
	}

	log.Println("✅ Conexión a MySQL exitosa")
	return db, nil
}
