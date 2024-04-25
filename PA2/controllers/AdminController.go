package controllers

import (
	"context"
	"pa2/models/dto"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/implements"
	"pa2/utils"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var (
	AdminRepo = implements.NewAdminRepository()
)

func AdminCreate(c *fiber.Ctx) error {
	InputReq := new(dto.RequestAdminCreate)

	if err := c.BodyParser(InputReq); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(InputReq); err != nil {
		response.BadRequestResponse(c)
	}

	password, err := utils.GeneratePassword(InputReq.Password)
	if err != nil {
		response.InternalServerError(c)
	}

	InputReq.Password = password

	admin := entity.Admin{
		Username: InputReq.Username,
		Password: password,
	}

	err = AdminRepo.AdminCreate(context.Background(), &admin)

	if err != nil {
		response.InternalServerError(c)
	}

	return c.JSON(admin)
}

func AdminLogin(c *fiber.Ctx) error {
	var credentials dto.RequestAdminLogin
	if err := c.BodyParser(&credentials); err != nil {
		return response.BadRequestResponse(c)
	}

	token, err := AdminRepo.AdminLogin(context.Background(), credentials)
	if err != nil {
		return response.UnauthorizedResponse(c)
	}

	cookie := fiber.Cookie{
		Name:     "admin",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}

func AdminGetProfile(c *fiber.Ctx) error {
	admin, err := AdminRepo.AdminGetProfile(c)
	if err != nil {
		return response.BadRequestResponse(c)
	}

	return c.JSON(fiber.Map{
		"admin": admin,
	})
}

func AdminLogout(c *fiber.Ctx) error {
	AdminRepo.AdminLogout(c)
	return c.JSON(fiber.Map{
		"message": "Logged out Successfully",
	})
}
