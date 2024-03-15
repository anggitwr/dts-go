package model

type Item struct {
	ItemId       int    `json:"item_id" gorm:"primaryKey"`
	ItemCode     string `json:"item_code"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity"`
	OrderOrderId int    `json:"order_id"`
}
