package controllers

import (
	"os"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/implements"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	ProductRepo      = implements.NewProductRepository()
	PathImageProduct = "./Public/file_product"
)

func init() {
	if _, err := os.Stat(PathImageProduct); os.IsNotExist(err) {
		os.Mkdir(PathImageProduct, os.ModePerm)
	}
}

func CreateProduct(c *fiber.Ctx) error {
	if err := ProductRepo.CreateProduct(c); err != nil {
		return response.InternalServerError(c)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "created successfully",
	})
}

func GetAllProduct(c *fiber.Ctx) error {
	var product []entity.Product

	err := ProductRepo.GetAllProduct(&product)
	if err != nil {
		return response.InternalServerError(c)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   product,
	})
}

func GetProductById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}

	var product entity.Product
	if err := ProductRepo.GetProductById(&product, productId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.NotFoundResponse(c)
		}
		return response.InternalServerError(c)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   product,
	})

}

func UpdateProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}

	var product entity.Product
	if err := ProductRepo.UpdateProduct(c, &product, productId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.NotFoundResponse(c)
		}
		return response.InternalServerError(c)
	}

	admin := c.Locals("admin").(entity.Admin)
	if product.AdminID != admin.Id {
		return c.Status(403).JSON(fiber.Map{
			"status":  "Failed",
			"message": "You are not allowed to update this product",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}
	var product entity.Product

	if err := ProductRepo.DeleteProduct(c, &product, productId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.NotFoundResponse(c)
		}
		return err
	}

	admin := c.Locals("admin").(entity.Admin)
	if product.AdminID != admin.Id {
		return c.Status(403).JSON(fiber.Map{
			"status":  "Failed",
			"message": "You are not allowed to update this product",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}
