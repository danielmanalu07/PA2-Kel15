package controllers

import (
	"fmt"
	"os"
	"pa2/database"
	"pa2/models"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const cashierPhoto = "./Photo"

func init() {
	if _, err := os.Stat(cashierPhoto); os.IsNotExist(err) {
		os.Mkdir(cashierPhoto, 0755)
	}
}

func IndexCashier(ctx *fiber.Ctx) error {
	var cashier []models.Cashier

	database.DB.Find(&cashier)

	if len(cashier) == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "Cashier not found",
		})
	}
	return ctx.JSON(cashier)
}

func CreateCashier(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	if username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please enter a username",
		})
	}

	passwordInput := ctx.FormValue("password")
	if passwordInput == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please enter a password",
		})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(passwordInput), 10)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	phone := ctx.FormValue("phone")
	if phone == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please enter a phone number",
		})
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Photo is required",
		})
	}

	filename := fmt.Sprintf("Photo_%s%s", username, filepath.Ext(photo.Filename))

	if err := ctx.SaveFile(photo, filepath.Join(cashierPhoto, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save photo",
		})
	}

	adminID := ctx.Locals("id").(string)

	var admin models.Admin
	database.DB.Where("id = ?", adminID).Find(&admin)

	cashier := models.Cashier{
		Username: username,
		Password: string(password),
		Phone:    phone,
		Photo:    filename,
		AdminID:  admin.Id,
	}

	database.DB.Create(&cashier)

	return ctx.JSON(cashier)
}

func ShowCashier(ctx *fiber.Ctx) error {
	cashierIDStr := ctx.Params("id")

	cashierID, err := strconv.Atoi(cashierIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}
	var cashier models.Cashier

	database.DB.Where("id = ?", cashierID).First(&cashier)

	return ctx.JSON(cashier)
}

func UpdateCashier(ctx *fiber.Ctx) error {
	cashierIDStr := ctx.Params("id")
	cashierID, err := strconv.Atoi(cashierIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var cashier models.Cashier
	database.DB.Where("id = ?", cashierID).First(&cashier)
	if cashierID != int(cashier.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cashier Not Found",
		})
	}

	username := ctx.FormValue("username")
	if username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username is required",
		})
	}
	phone := ctx.FormValue("phone")
	if phone == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "phone is required",
		})
	}

	newPhoto, err := ctx.FormFile("photo")
	if err == nil {
		if cashier.Photo != "" {
			oldPhotoPath := filepath.Join(cashierPhoto, cashier.Photo)
			os.Remove(oldPhotoPath)
		}

		filename := fmt.Sprintf("Photo_%d%s", cashierID, filepath.Ext(newPhoto.Filename))
		newPhotoPath := filepath.Join(cashierPhoto, filename)
		if err := ctx.SaveFile(newPhoto, newPhotoPath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save file photo",
			})
		}

		cashier.Photo = filename
	} else if !os.IsNotExist(err) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error processing photo",
		})
	}

	cashier.Username = username
	cashier.Phone = phone

	database.DB.Save(&cashier)
	return ctx.JSON(cashier)
}
func DeleteCashier(ctx *fiber.Ctx) error {
	cashierIDStr := ctx.Params("id")
	cashierID, err := strconv.Atoi(cashierIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var cashier models.Cashier
	database.DB.Where("id = ?", cashierID).First(&cashier)

	if cashierID != int(cashier.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cashier Not Found",
		})
	}

	if err := database.DB.Delete(&cashier).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete cashier",
		})
	}

	if cashier.Photo != "" {
		photoPath := filepath.Join(cashierPhoto, cashier.Photo)
		if err := os.Remove(photoPath); err != nil {
			fmt.Printf("Failed to delete photo file: %v\n", err)
		}
	}

	return ctx.JSON(fiber.Map{"message": "Cashier deleted successfully"})
}
