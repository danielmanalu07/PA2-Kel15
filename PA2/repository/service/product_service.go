package service

import (
	"pa2/models/entity"

	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetAllProduct(product *[]entity.Product) error
	GetProductById(product *entity.Product, id int) error
	UpdateProduct(ctx *fiber.Ctx, product *entity.Product, id int) error
	DeleteProduct(ctx *fiber.Ctx, product *entity.Product, id int) error
}
