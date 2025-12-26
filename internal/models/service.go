package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"` // em minutos
	Price       float64 `json:"price"`
}