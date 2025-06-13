package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv carga el archivo .env en las variables de entorno del sistema
func LoadEnv() {
	// Resuelve la ruta absoluta desde la raíz del proyecto
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("No se pudo resolver el path absoluto: %v", err)
	}

	envPath := filepath.Join(rootPath, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error al cargar .env: %v", err)
	}
}

// GetEnv retorna una variable de entorno o un valor por defecto
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
