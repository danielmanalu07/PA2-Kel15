package controllers

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	cartService service.CartService
}

func NewCartController(cs service.CartService) *CartController {
	return &CartController{cartService: cs}
}

func (cs *CartController) AddItemCart(ctx *fiber.Ctx) error {
	var input dto.RequestCartCreate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	cart, err := cs.cartService.AddItemToCart(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Couldn't add item cart")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": cart,
	})
}

func (cs *CartController) GetItemMyCart(ctx *fiber.Ctx) error {
	customer := ctx.Locals("customer").(entity.Customer)
	carts, err := cs.cartService.CartGetItemMyCart(customer.Id)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": carts,
	})
}

func (cs *CartController) DeleteMyCart(ctx *fiber.Ctx) error {
	CartID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}
	customer := ctx.Locals("customer").(entity.Customer)
	var cart entity.Cart
	err = cs.cartService.DeleteMyCart(customer.Id, uint(CartID))
	if cart.CustomerID != customer.Id && cart.Id != uint(CartID) {
		if err != nil {
			return utils.MessageJSON(ctx, 400, "Failed", "Couldn't delete items in cart")
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Cart items deleted successfully",
	})
}

func (cs *CartController) UpdateQuantity(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	var input dto.RequestCartUpdate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	customer := ctx.Locals("customer").(entity.Customer)

	cart, err := cs.cartService.UpdateMyCart(ctx, customer.Id, uint(Id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": cart,
	})
}
