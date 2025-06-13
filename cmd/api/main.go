package main

import (
	"github.com/D2JS-Corp/GrupdiBack/internal/config"
	"github.com/D2JS-Corp/GrupdiBack/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
