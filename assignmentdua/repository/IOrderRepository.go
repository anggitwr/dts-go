package repository

import "assignmentdua/model"

type IOrderRepository interface {
	Create(newOrder model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	Delete(order_id int) error
	Update(orderID int, updatedOrder model.Order) error
}
