package controllers

import (
	"fmt"
	"os"
	"pa2/database"
	"pa2/models"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const ImageTabel = "./FileTable"

func init() {
	if _, err := os.Stat(ImageTabel); os.IsNotExist(err) {
		os.Mkdir(ImageTabel, os.ModePerm)
	}
}

func CreateTable(ctx *fiber.Ctx) error {
	number := ctx.FormValue("number")
	if number == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Number is required",
		})
	}

	numberStr, err := strconv.Atoi(number)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Format Number",
		})
	}

	capacity := ctx.FormValue("capacity")
	if capacity == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Capacity is required",
		})
	}

	capacityStr, err := strconv.Atoi(capacity)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Format Capacity",
		})
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image is required",
		})
	}

	filename := fmt.Sprintf("Table_%d%s", numberStr, filepath.Ext(image.Filename))

	if err := ctx.SaveFile(image, filepath.Join(ImageTabel, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save image",
		})
	}

	adminID := ctx.Locals("id").(string)

	var admin models.Admin
	database.DB.Where("id = ?", adminID).Find(&admin)

	table := models.Table{
		Number:   int64(numberStr),
		Capacity: int64(capacityStr),
		Image:    filename,
		AdminID:  admin.Id,
	}

	database.DB.Create(&table)

	return ctx.JSON(table)
}

func IndexTable(ctx *fiber.Ctx) error {
	var table []models.Table

	database.DB.Find(&table)

	if len(table) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Table not found",
		})
	}

	return ctx.JSON(table)
}

func ShowTable(ctx *fiber.Ctx) error {
	TabelIDStr := ctx.Params("id")

	tableID, err := strconv.Atoi(TabelIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var table models.Table

	database.DB.Where("id = ?", tableID).First(&table)

	if tableID != int(table.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Table Not Found",
		})
	}
	return ctx.JSON(table)
}

func UpdateTable(ctx *fiber.Ctx) error {
	TabelIDStr := ctx.Params("id")

	tableID, err := strconv.Atoi(TabelIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}
	var table models.Table

	database.DB.Where("id = ?", tableID).First(&table)

	if tableID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Table Not Found",
		})
	}

	number := ctx.FormValue("number")
	if number == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Number is required",
		})
	}

	numberStr, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Format Number",
		})
	}

	capacity := ctx.FormValue("capacity")
	if capacity == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Capacity is required",
		})
	}

	capacityStr, err := strconv.ParseInt(capacity, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Format Capacity",
		})
	}

	newImage, err := ctx.FormFile("image")
	if err == nil {
		if table.Image == "" {
			oldImagePath := filepath.Join(ImageTabel, table.Image)
			os.Remove(oldImagePath)
		}

		filename := fmt.Sprintf("Table_%d%s", numberStr, filepath.Ext(newImage.Filename))
		newImagePath := filepath.Join(ImageTabel, filename)
		if err := ctx.SaveFile(newImage, newImagePath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save image",
			})
		}

		table.Image = filename
	} else if !os.IsNotExist(err) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Proccessing Image",
		})
	}

	result := database.DB.Model(&table).Updates(models.Table{Number: numberStr, Capacity: capacityStr})
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Updating",
			"error":   result.Error.Error(),
		})
	}
	return ctx.JSON(table)
}

func DeleteTable(ctx *fiber.Ctx) error {
	TabelIDStr := ctx.Params("id")

	tableID, err := strconv.Atoi(TabelIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}
	var table models.Table

	database.DB.Where("id = ?", tableID).First(&table)

	if tableID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Table Not Found",
		})
	}

	if err := database.DB.Delete(&table).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Delete Table",
		})
	}

	if table.Image != "" {
		imagePath := filepath.Join(ImageTabel, table.Image)
		if err := os.Remove(imagePath); err != nil {
			fmt.Printf("Failde to delete image File : %v\n", err)
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "Table deleted successfully",
	})
}
