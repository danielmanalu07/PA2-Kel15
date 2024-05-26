package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type CartRepository interface {
	AddItem(cart entity.Cart) (*entity.Cart, error)
	GetItemMyCart(id uint) ([]entity.Cart, error)
	Delete(customerId uint, id uint) error
	Update(customerId uint, id uint, cart *entity.Cart) (*entity.Cart, error)
}

type cartRepository struct{}

func (c *cartRepository) Update(customerId uint, id uint, cart *entity.Cart) (*entity.Cart, error) {
	if err := database.DB.Where("customer_id = ? AND id = ?", customerId, id).Save(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *cartRepository) Delete(customerId uint, id uint) error {
	if err := database.DB.Where("customer_id = ? AND id = ?", customerId, id).Delete(&entity.Cart{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) GetItemMyCart(id uint) ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := database.DB.Where("customer_id = ?", id).Find(&carts).Error; err != nil {
		return nil, err
	}

	return carts, nil
}

func (c *cartRepository) AddItem(cart entity.Cart) (*entity.Cart, error) {
	if err := database.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}
