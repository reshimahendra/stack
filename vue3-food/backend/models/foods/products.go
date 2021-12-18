package models

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    ID          string      `json:"id"`
    Sku         string      `json:"sku" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Price       float64     `json:"price" binding:"required,number"`
    IsReady     bool        `json:"is_ready,default=true"`
    IsAvailable bool        `json:"is_available,default=true"`
    Picture     string      `json:"picture"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

