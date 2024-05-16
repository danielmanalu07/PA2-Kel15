package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"

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

	admin, err := c.adminService.AdminLogin(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, fiber.StatusUnauthorized, "Failed", err.Error())
	}

	return ctx.JSON(admin)
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
