package model

import (
	"time"
)

type Order struct {
	OrderId      int       `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	CreatedAt    time.Time `json:"ordered_at"`
	Item         []Item    `json:"item"`
}
