package implements

import (
	"pa2/database"
	"pa2/models/dto"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryImpl struct{}

// DeleteCategory implements service.CategoryService.
func (c *CategoryImpl) DeleteCategory(ctx *fiber.Ctx, category *entity.Category, id int) error {
	err := database.DB.Debug().First(&category, "id = ?", id).Error
	if err != nil {
		return response.NotFoundResponse(ctx)
	}

	if err := database.DB.Debug().Delete(&category).Error; err != nil {
		return response.InternalServerError(ctx)
	}

	return err
}

func (c *CategoryImpl) UpdateCategory(ctx *fiber.Ctx, category *entity.Category, id int) error {
	inputReq := new(dto.RequestUpdateCategory)
	if err := ctx.BodyParser(inputReq); err != nil {
		return response.BadRequestResponse(ctx)
	}
	data := database.DB.First(&category, "id = ?", id)
	if data == nil {
		return response.NotFoundResponse(ctx)
	}

	if inputReq.Name != "" {
		category.Name = inputReq.Name
	}

	if inputReq.Description != "" {
		category.Description = inputReq.Description
	}

	result := database.DB.Save(category)

	if result.Error == nil {
		return response.InternalServerError(ctx)
	}
	return result.Error
}

func (c *CategoryImpl) GetCategoryById(ctx *fiber.Ctx, category *entity.Category, id int) error {
	result := database.DB.First(&category, "id = ?", id)
	if result == nil {
		return response.NotFoundResponse(ctx)
	}
	return result.Error
}

func (c *CategoryImpl) GetAllCategories(ctx *fiber.Ctx, categories *[]entity.Category) error {
	result := database.DB.Find(&categories)
	if result == nil {
		return response.NotFoundResponse(ctx)
	}
	return result.Error
}

func (c *CategoryImpl) CreateCategory(ctx *fiber.Ctx, category *entity.Category) error {
	result := database.DB.Create(&category)
	return result.Error
}

func NewCategoryRepository() service.CategoryService {
	return &CategoryImpl{}
}
