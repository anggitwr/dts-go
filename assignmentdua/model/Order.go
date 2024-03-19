package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderId      int       `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	CreatedAt    time.Time `json:"ordered_at"`
	Item         []Item    `json:"item"`
	DeletedAt    gorm.DeletedAt
}
