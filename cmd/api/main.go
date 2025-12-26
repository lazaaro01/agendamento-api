package main

import (
	"agendamento-api/internal/config"
	"agendamento-api/internal/database"
	"agendamento-api/internal/models"
	"agendamento-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}