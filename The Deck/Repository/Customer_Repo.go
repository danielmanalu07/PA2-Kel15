package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	utils "api/the_deck/Utils"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CustomerRepository interface {
	Register(customer entity.Customer) (*entity.Customer, error)
	Login(input dto.RequestCustomerLogin) (*entity.Customer, error)
	Profile(customer entity.Customer) (*entity.Customer, error)
	Logout() (*fiber.Cookie, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	EditPW(customer *entity.Customer) (*entity.Customer, error)
	Forgot(customer *entity.Customer) (*entity.Customer, error)
}

type customerRepository struct{}

func (c *customerRepository) Forgot(customer *entity.Customer) (*entity.Customer, error) {
	if err := database.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerRepository) EditPW(customer *entity.Customer) (*entity.Customer, error) {
	if err := database.DB.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *customerRepository) Update(customer *entity.Customer) (*entity.Customer, error) {
	if err := database.DB.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *customerRepository) Logout() (*fiber.Cookie, error) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	return &cookie, nil
}

func (c *customerRepository) Profile(customer entity.Customer) (*entity.Customer, error) {
	return &customer, nil
}

func (c *customerRepository) Login(input dto.RequestCustomerLogin) (*entity.Customer, error) {
	var customer entity.Customer
	result := database.DB.First(&customer, "email = ?", input.Email)
	if result.Error != nil {
		return nil, result.Error
	}

	checkPw := utils.CheckPassword(input.Password, customer.Password)
	if !checkPw {
		return nil, errors.New("Inccorrect Password")
	}

	return &customer, nil
}

func (c *customerRepository) Register(customer entity.Customer) (*entity.Customer, error) {
	if err := database.DB.Create(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}
