package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type OrderRepository interface {
	Create(order entity.Order) (*entity.Order, error)
	GetAll() ([]entity.Order, error)
	GetById(id uint) (*entity.Order, error)
	GetMyOrder(customerId uint) ([]entity.Order, error)
	GetMyOrderById(customerId uint, id uint) (*entity.Order, error)
	Payment(customerId uint, id uint, order *entity.Order) (*entity.Order, error)
	Update(order *entity.Order) (*entity.Order, error)
}

type orderRepository struct{}

func (o *orderRepository) GetById(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := database.DB.First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *orderRepository) Update(order *entity.Order) (*entity.Order, error) {
	if err := database.DB.Save(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderRepository) GetMyOrderById(customerId uint, id uint) (*entity.Order, error) {
	var orders entity.Order
	if err := database.DB.Where("customer_id = ? AND id = ?", customerId, id).First(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

func (o *orderRepository) Payment(customerId uint, id uint, order *entity.Order) (*entity.Order, error) {
	if err := database.DB.Where("customer_id = ? AND id = ?", customerId, id).Save(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderRepository) GetMyOrder(customerId uint) ([]entity.Order, error) {
	var orders []entity.Order
	if err := database.DB.Where("customer_id = ?", customerId).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderRepository) GetAll() ([]entity.Order, error) {
	var order []entity.Order
	if err := database.DB.Preload("Products").Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderRepository) Create(order entity.Order) (*entity.Order, error) {
	if err := database.DB.Debug().Preload("Tables").Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
