package routes

import (
	"agendamento-api/internal/handlers"
	"agendamento-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes (JWT)
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/services", handlers.CreateService)
		protected.GET("/services", handlers.ListServices)

		protected.POST("/appointments", handlers.CreateAppointment)
		protected.GET("/appointments", handlers.ListAppointments)
		protected.PUT("/appointments/:id/cancel", handlers.CancelAppointment)
	}
}