package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(cs service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: cs}
}

func (c *CategoryController) CategoryCreate(ctx *fiber.Ctx) error {
	var input dto.RequestCategoryCreate
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Invalid request payload")
	}
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	category, err := c.categoryService.CategoryCreate(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Cannot create category")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func (c *CategoryController) CategoryGetAll(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.CategoryGetAll()
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": categories,
	})
}

func (c *CategoryController) CategoryGetById(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}
	category, err := c.categoryService.CategoryGetById(uint(Id))
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func (c *CategoryController) CategoryUpdate(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestCategoryUpdate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	category, err := c.categoryService.CategoryUpdate(ctx, uint(Id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func (c *CategoryController) CategoryDelete(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	if err := c.categoryService.CategoryDelete(uint(Id)); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Category Successfully",
	})
}
