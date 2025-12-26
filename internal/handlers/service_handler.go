package handlers

import (
	"net/http"

	"agendamento-api/internal/database"
	"agendamento-api/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := database.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar serviço"})
		return
	}

	c.JSON(http.StatusCreated, service)
}

func ListServices(c *gin.Context) {
	var services []models.Service
	database.DB.Find(&services)
	c.JSON(http.StatusOK, services)
}