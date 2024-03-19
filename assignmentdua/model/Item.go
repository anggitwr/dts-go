package model

import "gorm.io/gorm"

type Item struct {
	ItemId       int    `json:"item_id" gorm:"primaryKey"`
	ItemCode     string `json:"item_code"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity"`
	OrderOrderId int
	DeletedAt    gorm.DeletedAt
}
