package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(ps service.ProductService) *ProductController {
	return &ProductController{productService: ps}
}

func (p *ProductController) ProductCreate(ctx *fiber.Ctx) error {
	var input dto.RequestProductCreate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Invalid request payload")
	}
	product, err := p.productService.ProductCreate(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Couldn't create product")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func (p *ProductController) ProductGetAll(ctx *fiber.Ctx) error {
	products, err := p.productService.ProductGetAll()
	if err != nil {
		return utils.MessageJSON(ctx, 500, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": products,
	})
}

func (p *ProductController) ProductGetById(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	product, err := p.productService.ProductGetById(uint(Id))
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func (p *ProductController) ProductUpdate(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestProductUpdate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	product, err := p.productService.ProductUpdate(ctx, uint(Id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func (p *ProductController) ProductDelete(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	if err := p.productService.ProductDelete(uint(Id)); err != nil {
		if err.Error() == "record not found" {
			return utils.MessageJSON(ctx, 404, "Failed", err.Error())
		}
		return utils.MessageJSON(ctx, 500, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Product Successfully",
	})
}
