package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(cs service.CustomerService) *CustomerController {
	return &CustomerController{customerService: cs}
}

func (cs *CustomerController) CustomerRegister(ctx *fiber.Ctx) error {
	var input dto.RequestCustomerRegister
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Invalid Request")
	}
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	customer, err := cs.customerService.CustomerRegister(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 500, "Failed", "Cannot Register Customer")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func (cs *CustomerController) CustomerLogin(ctx *fiber.Ctx) error {
	input := new(dto.RequestCustomerLogin)
	if err := ctx.BodyParser(input); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	customer, token, err := cs.customerService.CustomerLogin(ctx, *input)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
		"token":   token,
	})
}

func (cs *CustomerController) GetProfile(ctx *fiber.Ctx) error {
	customer, err := cs.customerService.GetProfile(ctx)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", "Unauthenticated")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func (cs *CustomerController) CustomerLogout(ctx *fiber.Ctx) error {
	cookie, err := cs.customerService.CustomerLogout(ctx)
	if err != nil {
		return err
	}

	ctx.Cookie(cookie)

	ctx.Locals("customer", nil)

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Successfully",
	})
}

func (cs *CustomerController) CustomerUpdateProfile(ctx *fiber.Ctx) error {
	var input dto.RequestCustomerUpdateProfile
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	customer, err := cs.customerService.CustomerUpdate(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func (cs *CustomerController) CustomerEditPassword(ctx *fiber.Ctx) error {
	var input dto.RequestCustomerEditPassword
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	customer, err := cs.customerService.CustomerEditPassword(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer.Password,
	})
}

func (cs *CustomerController) CustomerForgotPassword(ctx *fiber.Ctx) error {
	var input dto.RequestCustomerForgotPassword
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	customer, err := cs.customerService.CustomerForgotPassword(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer.Password,
	})
}
