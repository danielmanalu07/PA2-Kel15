package controllers

import (
	"pa2/models/dto"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/implements"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	CategoryRepo = implements.NewCategoryRepository()
)

func CreateCategory(c *fiber.Ctx) error {
	categoryReq := new(dto.RequestCreateCategory)
	if err := c.BodyParser(categoryReq); err != nil {
		return response.BadRequestResponse(c)
	}

	validate := validator.New()
	if err := validate.Struct(categoryReq); err != nil {
		return response.BadRequestResponse(c)
	}

	Admin := c.Locals("admin").(entity.Admin)

	category := entity.Category{
		Name:        categoryReq.Name,
		Description: categoryReq.Description,
		AdminID:     Admin.Id,
	}

	if err := CategoryRepo.CreateCategory(c, &category); err != nil {
		return response.InternalServerError(c)
	}
	return c.Status(fiber.StatusCreated).JSON(category)
}

func GetAllCategory(c *fiber.Ctx) error {
	var category []entity.Category

	err := CategoryRepo.GetAllCategories(c, &category)
	if err != nil {
		response.InternalServerError(c)
	}
	if len(category) == 0 {
		return response.NotFoundResponse(c)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   category,
	})
}

func GetAllCategoryById(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}

	var category entity.Category
	if err := CategoryRepo.GetCategoryById(c, &category, categoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.NotFoundResponse(c)
		}
		return response.InternalServerError(c)
	}
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}
	var category entity.Category
	if err := CategoryRepo.UpdateCategory(c, &category, categoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.NotFoundResponse(c)
		}
		return response.InternalServerError(c)
	}
	admin := c.Locals("admin").(entity.Admin)
	if category.AdminID != admin.Id {
		return c.Status(403).JSON(fiber.Map{
			"status":  "Failed",
			"message": "You are not allowed to update this category",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequestResponse(c)
	}

	var category entity.Category

	CategoryRepo.DeleteCategory(c, &category, categoryId)

	admin := c.Locals("admin").(entity.Admin)
	if category.AdminID != admin.Id {
		return c.Status(403).JSON(fiber.Map{
			"status":  "Failed",
			"message": "You are not allowed to delete this category",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Category deleted successfully",
	})
}
