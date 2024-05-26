package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartService interface {
	AddItemToCart(ctx *fiber.Ctx, input dto.RequestCartCreate) (*response.CartResponse, error)
	CartGetItemMyCart(id uint) ([]response.CartResponse, error)
	DeleteMyCart(CustomerID uint, id uint) error
	UpdateMyCart(ctx *fiber.Ctx, CustomerID uint, id uint, input dto.RequestCartUpdate) (*response.CartResponse, error)
}

type cartService struct {
	cartService repository.CartRepository
}

func (c *cartService) UpdateMyCart(ctx *fiber.Ctx, CustomerID uint, id uint, input dto.RequestCartUpdate) (*response.CartResponse, error) {
	// Check if the item exists in the cart
	cartItem, err := c.cartService.GetItemByID(CustomerID, id)
	if err != nil {
		return nil, errors.New("item not found in the cart")
	}

	if input.Quantity != 0 {
		cartItem.Quantity = input.Quantity
	}

	updatedCart, err := c.cartService.Update(CustomerID, id, cartItem)
	if err != nil {
		return nil, err
	}

	cartResponse := &response.CartResponse{
		Id:         updatedCart.Id,
		ProductID:  updatedCart.ProductID,
		CustomerID: updatedCart.CustomerID,
		Quantity:   updatedCart.Quantity,
	}

	return cartResponse, nil
}
func (c *cartService) DeleteMyCart(CustomerID uint, id uint) error {
	result := c.cartService.Delete(CustomerID, id)
	if result != nil {
		return errors.New("Cannot delete Cart Item")
	}
	return result
}

func (c *cartService) CartGetItemMyCart(id uint) ([]response.CartResponse, error) {
	cart, err := c.cartService.GetItemMyCart(id)
	if err != nil {
		return nil, err
	}

	var cartResponses []response.CartResponse
	for _, carts := range cart {
		cartResponse := response.CartResponse{
			Id:         carts.Id,
			CustomerID: carts.CustomerID,
			ProductID:  carts.ProductID,
			Quantity:   carts.Quantity,
		}

		cartResponses = append(cartResponses, cartResponse)
	}

	return cartResponses, nil
}

func (c *cartService) AddItemToCart(ctx *fiber.Ctx, input dto.RequestCartCreate) (*response.CartResponse, error) {
	ProductId, err := strconv.Atoi(ctx.FormValue("product_id"))
	if err != nil {
		return nil, err
	}

	input.ProductID = uint(ProductId)

	customer := ctx.Locals("customer").(entity.Customer)

	cart := entity.Cart{
		CustomerID: customer.Id,
		Quantity:   1,
		ProductID:  uint(ProductId),
	}

	createCart, err := c.cartService.AddItem(cart)
	if err != nil {
		return nil, err
	}

	response := &response.CartResponse{
		Id:         createCart.Id,
		CustomerID: createCart.CustomerID,
		ProductID:  createCart.ProductID,
		Quantity:   createCart.Quantity,
	}

	return response, nil
}

func NewCartService(cr repository.CartRepository) CartService {
	return &cartService{cartService: cr}
}
