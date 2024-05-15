package controllers

import (
	"fmt"
	"os"
	"pa2/database"
	"pa2/models"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const PaymentPhoto = "./PaymentFile"

func init() {
	if _, err := os.Stat(PaymentPhoto); os.IsNotExist(err) {
		os.Mkdir(PaymentPhoto, os.ModePerm)
	}
}

func SendPayment(ctx *fiber.Ctx) error {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Photo is required",
		})
	}

	customerID := ctx.Locals("id").(string)

	var customer models.Customer
	database.DB.Where("id = ?", customerID).Find(&customer)

	filename := fmt.Sprintf("Payment_%d%s", customer.ID, filepath.Ext(photo.Filename))

	if err := ctx.SaveFile(photo, filepath.Join(PaymentPhoto, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed To Save Payment",
		})
	}

	payments := models.ProofOfPayment{
		Photo:      filename,
		CustomerID: customer.ID,
	}
	database.DB.Create(&payments)
	return ctx.JSON(payments)
}
