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

const ImageProduct = "./FileProduct"

func init() {
	if _, err := os.Stat(ImageProduct); os.IsNotExist(err) {
		os.Mkdir(ImageProduct, os.ModePerm)
	}
}

func CreateProduct(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "name is required",
		})
	}

	description := ctx.FormValue("description")
	if description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "description is required",
		})
	}

	price := ctx.FormValue("price")
	if price == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "price is required",
		})
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image is required",
		})
	}

	categoryID := ctx.FormValue("category_id")
	if categoryID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "category_id is required",
		})
	}

	categoryIDStr, err := strconv.Atoi(categoryID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid format ID",
		})
	}

	filename := fmt.Sprintf("Product_%s%s", name, filepath.Ext(image.Filename))

	if err := ctx.SaveFile(image, filepath.Join(ImageProduct, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed To Save File Product",
		})
	}

	adminId := ctx.Locals("id").(string)

	var admin models.Admin
	database.DB.Where("id = ?", adminId).First(&admin)

	categoryId := ctx.Locals("id").(string)

	var category models.Category
	database.DB.Where("id = ?", categoryId).First(&category)

	products := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Image:       filename,
		AdminID:     admin.Id,
		CategoryID:  uint(categoryIDStr),
	}

	database.DB.Create(&products)

	return ctx.JSON(products)
}

func IndexProduct(ctx *fiber.Ctx) error {
	var product []models.Product

	database.DB.Find(&product)

	if len(product) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func ShowProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")

	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product

	database.DB.Where("id = ?", productID).First(&product)

	if productID != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")

	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product

	database.DB.Where("id = ?", productID).First(&product)

	if productID != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "name is required",
		})
	}

	description := ctx.FormValue("description")
	if description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "description is required",
		})
	}

	price := ctx.FormValue("price")
	if price == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "price is required",
		})
	}

	categoryIDStr := ctx.FormValue("category_id")
	if categoryIDStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "category_id is required",
		})
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid format ID",
		})
	}

	newImage, err := ctx.FormFile("image")
	if err == nil {
		if product.Image == "" {
			oldImagePath := filepath.Join(ImageProduct, product.Image)
			os.Remove(oldImagePath)
		}

		filename := fmt.Sprintf("Product_%s%s", name, filepath.Ext(newImage.Filename))
		newImagePath := filepath.Join(ImageProduct, filename)
		if err := ctx.SaveFile(newImage, newImagePath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save image",
			})
		}

		product.Image = filename
	} else if !os.IsNotExist(err) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Proccessing Image",
		})
	}

	result := database.DB.Model(&product).Updates(models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  uint(categoryID),
	})
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Updating",
			"error":   result.Error.Error(),
		})
	}

	return ctx.JSON(product)

}

func DeleteProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")

	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product

	database.DB.Where("id = ?", productID).First(&product)

	if productID != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Delete Table",
		})
	}

	if product.Image != "" {
		imagePath := filepath.Join(ImageProduct, product.Image)
		if err := os.Remove(imagePath); err != nil {
			fmt.Printf("Failde to delete image File : %v\n", err)
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
