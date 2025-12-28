package handlers

import (
	"log"
	"net/http"
	"time"

	"agendamento-api/internal/database"
	"agendamento-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AppointmentInput struct {
	ServiceID uint      `json:"service_id"`
	StartTime time.Time `json:"start_time"`
}

func CreateAppointment(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var input AppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Erro no bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	log.Printf("Recebendo agendamento: UserID=%d, ServiceID=%d, StartTime=%v", userID, input.ServiceID, input.StartTime)

	var service models.Service
	if err := database.DB.First(&service, input.ServiceID).Error; err != nil {
		log.Printf("Serviço %d não encontrado: %v", input.ServiceID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado"})
		return
	}

	start := input.StartTime
	end := start.Add(time.Minute * time.Duration(service.Duration))

	var count int64
	database.DB.Model(&models.Appointment{}).
		Where("start_time < ? AND end_time > ? AND status = ?", end, start, "CONFIRMED").
		Count(&count)

	if count > 0 {
		log.Printf("Conflito de horário detectado para o início %v", start)
		c.JSON(http.StatusConflict, gin.H{"error": "Horário indisponível"})
		return
	}

	appointment := models.Appointment{
		UserID:    userID,
		ServiceID: service.ID,
		StartTime: start,
		EndTime:   end,
		Status:    "CONFIRMED",
	}

	if err := database.DB.Create(&appointment).Error; err != nil {
		log.Printf("Erro ao criar agendamento no banco: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar agendamento"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func ListAppointments(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var appointments []models.Appointment
	database.DB.Preload("Service").Where("user_id = ?", userID).Find(&appointments)

	c.JSON(http.StatusOK, appointments)
}

func CancelAppointment(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Model(&models.Appointment{}).
		Where("id = ?", id).
		Update("status", "CANCELED")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar agendamento"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agendamento cancelado"})
}
