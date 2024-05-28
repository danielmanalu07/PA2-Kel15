package controllers

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

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

func (oc *OrderController) GetAllOrder(ctx *fiber.Ctx) error {
	orders, err := oc.orderService.GetAllOrder()
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": orders,
	})
}

func (oc *OrderController) GetMyOrder(ctx *fiber.Ctx) error {
	customer := ctx.Locals("customer").(entity.Customer)
	orders, err := oc.orderService.GetMyOrder(customer.Id)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": orders,
	})
}

func (oc *OrderController) ProofOfPayment(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	customer := ctx.Locals("customer").(entity.Customer)

	order, err := oc.orderService.ProofPayment(ctx, customer.Id, uint(id))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}

func (oc *OrderController) UpdateStatus(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestOrderUpdateStatus
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	order, err := oc.orderService.UpdateStatus(uint(id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}
