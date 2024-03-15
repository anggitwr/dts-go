package repository

import (
	"assignmentdua/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) Create(newOrder model.Order) (model.Order, error) {

	tx := or.db.Create(&newOrder)

	return newOrder, tx.Error
}

func (or *OrderRepository) GetAll() ([]model.Order, error) {
	var orders = []model.Order{}

	tx := or.db.Preload("Item").Find(&orders)

	return orders, tx.Error
}

func (or *OrderRepository) Delete(order_id int) error {
	// fmt.Println(uuid)
	tx := or.db.Unscoped().Delete(&model.Order{}, "order_id = ?", order_id)
	return tx.Error
}
func (or *OrderRepository) Update(order_id int, updatedOrder model.Order) error {

	var existingOrder model.Order
	result := or.db.First(&existingOrder, order_id)
	if result.Error != nil {
		return result.Error
	}

	existingOrder.CustomerName = updatedOrder.CustomerName

	tx := or.db.Save(&existingOrder)
	return tx.Error
}
