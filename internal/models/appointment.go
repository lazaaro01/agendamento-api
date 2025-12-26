package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	ServiceID uint      `json:"service_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}