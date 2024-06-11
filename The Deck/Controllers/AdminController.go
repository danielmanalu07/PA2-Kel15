package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	adminService service.AdminService
}

func NewAdminController(as service.AdminService) *AdminController {
	return &AdminController{adminService: as}
}

func (c *AdminController) AdminLogin(ctx *fiber.Ctx) error {
	var input dto.RequestAdminLogin
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, fiber.StatusBadRequest, "Failed", "Invalid request body")
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	admin, token, err := c.adminService.AdminLogin(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, fiber.StatusUnauthorized, "Failed", err.Error())
	}

	return ctx.JSON(fiber.Map{
		"status":  "message",
		"message": admin,
		"token":   token,
	})
}

func (c *AdminController) GetProfile(ctx *fiber.Ctx) error {
	admin, err := c.adminService.GetProfile(ctx)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", "Unauthenticated")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": admin,
	})
}

func (c *AdminController) LogoutAdmin(ctx *fiber.Ctx) error {
	cookie, err := c.adminService.LogoutAdmin(ctx)
	if err != nil {
		return err
	}

	ctx.Cookie(cookie)

	ctx.Locals("admin", nil)

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout successfully",
	})
}

func (c *AdminController) UpdateStatusOrder(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestOrderUpdateStatus
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	order, err := c.adminService.UpdateStatus(ctx, uint(id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}

func (c *AdminController) ApproveReqTable(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	var input dto.UpdateRequestTable
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	req_table, err := c.adminService.ApproveReqTable(ctx, uint(id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": req_table,
	})
}
