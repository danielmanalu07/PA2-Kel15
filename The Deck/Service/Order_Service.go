package service

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderService interface {
	CreateOrder(ctx *fiber.Ctx, input dto.RequestOrderCreate) (*response.OrderResponse, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func GenerateCodeOrder() string {
	var lastOrder entity.Order
	today := time.Now().Format("2006-01-02")
	database.DB.Where("DATE(created_at) = ?", today).Last(&lastOrder)

	var count int64
	database.DB.Model(&entity.Order{}).Where("DATE(created_at) = ?", today).Count(&count)

	return fmt.Sprintf("ORD-%03d", count+1)
}

func (o *orderService) CreateOrder(ctx *fiber.Ctx, input dto.RequestOrderCreate) (*response.OrderResponse, error) {
	// Validate PaymentMethod
	if input.PaymentMethod != "Cash" && input.PaymentMethod != "QRIS" {
		return nil, utils.MessageJSON(ctx, 400, "error", "Invalid Payment Method")
	}

	// Validate PickUpType and TableId
	if input.PickUpType == "Dine in" && input.TableId == 0 {
		return nil, utils.MessageJSON(ctx, 400, "error", "Table is required for Dine in")
	} else if input.PickUpType != "Take Away" && input.PickUpType != "Dine in" {
		return nil, utils.MessageJSON(ctx, 400, "error", "Invalid Pick Up Type")
	}

	// Get customer from context
	customer := ctx.Locals("customer").(entity.Customer)

	// Fetch cart items for the customer
	var cartItems []entity.Cart
	if err := database.DB.Where("customer_id = ?", customer.Id).Preload("Product").Find(&cartItems).Error; err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not fetch cart items")
	}

	// Create Order
	order := entity.Order{
		Code:           GenerateCodeOrder(),
		CustomerID:     customer.Id,
		Total:          input.Total,
		Note:           input.Note,
		PaymentMethod:  input.PaymentMethod,
		TableId:        input.TableId,
		PickUpType:     input.PickUpType,
		Status:         0,
		ProofOfPayment: "",
	}

	// Save the order
	savedOrder, err := o.orderRepository.Create(order)
	if err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not create order")
	}

	// Create order products
	for _, productID := range input.ProductIDs {
		orderProduct := entity.OrderProduct{
			OrderID:   savedOrder.Id,
			ProductID: productID,
		}
		if err := database.DB.Create(&orderProduct).Error; err != nil {
			return nil, utils.MessageJSON(ctx, 500, "error", "Could not associate product with order")
		}
	}

	// Retrieve the order with products
	if err := database.DB.Preload("Products").First(&savedOrder, savedOrder.Id).Error; err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not retrieve created order")
	}

	orderResponse := &response.OrderResponse{
		Id:             savedOrder.Id,
		Code:           savedOrder.Code,
		Products:       savedOrder.Products,
		Total:          savedOrder.Total,
		Note:           savedOrder.Note,
		PaymentMethod:  savedOrder.PaymentMethod,
		TableId:        savedOrder.TableId,
		PickUpType:     savedOrder.PickUpType,
		ProofOfPayment: savedOrder.ProofOfPayment,
		Status:         savedOrder.Status,
		CreatedAt:      savedOrder.CreatedAt,
		UpdatedAt:      savedOrder.UpdatedAt,
	}

	return orderResponse, nil
}

func NewOrderService(or repository.OrderRepository) OrderService {
	return &orderService{orderRepository: or}
}
