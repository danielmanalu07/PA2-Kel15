package service

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

const PathImageOrder = "./Public/Order"

type OrderService interface {
	CreateOrder(ctx *fiber.Ctx, input dto.RequestOrderCreate) (*response.OrderResponse, error)
	GetAllOrder() ([]response.OrderResponse, error)
	GetMyOrder(customerId uint) ([]response.OrderResponse, error)
	ProofPayment(ctx *fiber.Ctx, customerId uint, id uint) (*response.OrderResponse, error)
	UpdateStatus(id uint, input dto.RequestOrderUpdateStatus) (*response.OrderResponse, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func (o *orderService) UpdateStatus(id uint, input dto.RequestOrderUpdateStatus) (*response.OrderResponse, error) {
	order, err := o.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Status != 0 {
		order.Status = input.Status
	}

	save, err := o.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	response := &response.OrderResponse{
		Id:             save.Id,
		Code:           save.Code,
		CustomerID:     save.CustomerID,
		Products:       save.Products,
		Total:          save.Total,
		Note:           save.Note,
		PaymentMethod:  save.PaymentMethod,
		TableId:        save.TableId,
		PickUpType:     save.PickUpType,
		ProofOfPayment: save.ProofOfPayment,
		Status:         save.Status,
	}

	return response, nil
}

func (o *orderService) ProofPayment(ctx *fiber.Ctx, customerId uint, id uint) (*response.OrderResponse, error) {
	orders, err := o.orderRepository.GetMyOrderById(customerId, id)
	if err != nil {
		return nil, err
	}

	image, err := ctx.FormFile("image")
	if err == nil {
		if orders.ProofOfPayment != "" {
			oldPath := filepath.Join(PathImageOrder, orders.ProofOfPayment)
			os.Remove(oldPath)
		}

		newFilename := utils.GenerateImageFile(orders.Code, image.Filename)
		if err := ctx.SaveFile(image, filepath.Join(PathImageOrder, newFilename)); err != nil {
			return nil, err
		}

		orders.ProofOfPayment = newFilename
	}

	save, err := o.orderRepository.Payment(customerId, id, orders)
	if err != nil {
		return nil, err
	}

	orderResponse := &response.OrderResponse{
		Id:             save.Id,
		Code:           save.Code,
		CustomerID:     save.CustomerID,
		Products:       save.Products,
		Total:          save.Total,
		Note:           save.Note,
		PaymentMethod:  save.PaymentMethod,
		TableId:        save.TableId,
		PickUpType:     save.PickUpType,
		ProofOfPayment: save.ProofOfPayment,
		Status:         save.Status,
	}

	return orderResponse, nil

}

func (o *orderService) GetMyOrder(customerId uint) ([]response.OrderResponse, error) {
	order, err := o.orderRepository.GetMyOrder(customerId)
	if err != nil {
		return nil, err
	}

	var orderResponses []response.OrderResponse
	for _, orders := range order {
		orderResponse := response.OrderResponse{
			Id:             orders.Id,
			Code:           orders.Code,
			CustomerID:     orders.CustomerID,
			Products:       orders.Products,
			Total:          orders.Total,
			Note:           orders.Note,
			PaymentMethod:  orders.PaymentMethod,
			TableId:        orders.TableId,
			PickUpType:     orders.PickUpType,
			ProofOfPayment: orders.ProofOfPayment,
			Status:         orders.Status,
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (o *orderService) GetAllOrder() ([]response.OrderResponse, error) {
	orders, err := o.orderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var orderResponses []response.OrderResponse
	for _, order := range orders {
		orderResponse := response.OrderResponse{
			Id:             order.Id,
			Code:           order.Code,
			CustomerID:     order.CustomerID,
			Products:       order.Products,
			Total:          order.Total,
			Note:           order.Note,
			PaymentMethod:  order.PaymentMethod,
			TableId:        order.TableId,
			PickUpType:     order.PickUpType,
			ProofOfPayment: order.ProofOfPayment,
			Status:         order.Status,
			CreatedAt:      order.CreatedAt,
			UpdatedAt:      order.UpdatedAt,
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
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
	if input.PaymentMethod != "Cash" && input.PaymentMethod != "QRIS" {
		return nil, utils.MessageJSON(ctx, 400, "error", "Invalid Payment Method")
	}

	if input.PickUpType == "Dine In" && input.TableId == 0 {
		return nil, utils.MessageJSON(ctx, 400, "error", "Table is required for Dine In")
	} else if input.PickUpType != "Take Away" && input.PickUpType != "Dine In" {
		return nil, utils.MessageJSON(ctx, 400, "error", "Invalid Pick Up Type")
	}

	// Mendapatkan informasi pelanggan dari konteks
	customer, ok := ctx.Locals("customer").(entity.Customer)
	if !ok {
		return nil, utils.MessageJSON(ctx, 400, "error", "Customer not found")
	}

	// Mendapatkan item keranjang
	var cartItems []entity.Cart
	if err := database.DB.Where("customer_id = ?", customer.Id).Preload("Product").Find(&cartItems).Error; err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not fetch cart items")
	}

	if len(cartItems) == 0 {
		return nil, utils.MessageJSON(ctx, 400, "error", "Cart is empty")
	}

	// Set TableId to nil if PickUpType is "Take Away"
	var tableId *uint
	if input.PickUpType == "Take Away" {
		tableId = nil
	} else {
		tableId = &input.TableId
	}

	// Membuat pesanan
	order := entity.Order{
		Code:           GenerateCodeOrder(),
		CustomerID:     customer.Id,
		Total:          input.Total,
		Note:           input.Note,
		PaymentMethod:  input.PaymentMethod,
		TableId:        tableId,
		PickUpType:     input.PickUpType,
		Status:         0,
		ProofOfPayment: "",
	}

	savedOrder, err := o.orderRepository.Create(order)
	if err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not create order")
	}

	for _, cartItem := range cartItems {
		orderProduct := entity.OrderProduct{
			OrderID:   savedOrder.Id,
			ProductID: cartItem.ProductID,
		}
		if err := database.DB.Create(&orderProduct).Error; err != nil {
			return nil, utils.MessageJSON(ctx, 500, "error", "Could not associate product with order")
		}
	}

	if err := database.DB.Preload("Products").First(&savedOrder, savedOrder.Id).Error; err != nil {
		return nil, utils.MessageJSON(ctx, 500, "error", "Could not retrieve created order")
	}

	// Membuat respons pesanan
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
