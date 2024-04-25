package service

import (
	"pa2/models/entity"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	CreateCategory(ctx *fiber.Ctx, category *entity.Category) error
	GetAllCategories(ctx *fiber.Ctx, category *[]entity.Category) error
	GetCategoryById(ctx *fiber.Ctx, category *entity.Category, id int) error
	UpdateCategory(ctx *fiber.Ctx, category *entity.Category, id int) error
	DeleteCategory(ctx *fiber.Ctx, category *entity.Category, id int) error
}
