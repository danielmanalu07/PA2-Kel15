package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type OrderRepository interface {
	Create(order entity.Order) (*entity.Order, error)
}

type orderRepository struct{}

func (o *orderRepository) Create(order entity.Order) (*entity.Order, error) {
	if err := database.DB.Debug().Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
