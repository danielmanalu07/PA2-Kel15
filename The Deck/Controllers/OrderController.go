package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(os service.OrderService) *OrderController {
	return &OrderController{orderService: os}
}

func (oc *OrderController) CustomerCreateOrder(ctx *fiber.Ctx) error {
	var input dto.RequestOrderCreate
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Invalid request payload")
	}

	order, err := oc.orderService.CreateOrder(ctx, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}
