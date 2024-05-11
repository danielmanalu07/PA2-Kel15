package controllers

import (
	database "cashier/Database"
	"cashier/Models/dto"
	"cashier/Models/entity"
	utils "cashier/Utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CashierRegister(c *fiber.Ctx) error {
	input := new(dto.RequestCashierRegister)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	password, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate password",
		})
	}

	input.Password = password

	cashier := entity.Cashier{
		Name:     input.Name,
		Phone:    input.Phone,
		Username: input.Username,
		Password: password,
	}

	result := database.DB.Create(&cashier)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't to create cashier account",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": cashier,
	})
}

func CashierLogin(c *fiber.Ctx) error {
	input := new(dto.RequestCashierLogin)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	var cashier entity.Cashier

	result := database.DB.First(&cashier, "username = ?", input.Username)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Username Not Found",
		})
	}

	checkpw := utils.CheckPassword(input.Password, cashier.Password)

	if !checkpw {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Inccorrect Password",
		})
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(cashier.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error Generating Token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": cashier,
		"token":   token,
	})
}

func CashierGetProfile(c *fiber.Ctx) error {
	cashier := c.Locals("cashier").(entity.Cashier)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": cashier,
	})
}

func CashierLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Locals("cashier", nil)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Successfully",
	})
}
