package controllers

import (
	"api/the_deck/Models/dto"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TableController struct {
	tableService service.TableService
}

func NewTableController(ts service.TableService) *TableController {
	return &TableController{tableService: ts}
}

func (t *TableController) TableCreate(ctx *fiber.Ctx) error {
	var input dto.RequestTableCreate
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}
	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return err
	}

	table, err := t.tableService.TableCreate(ctx, input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Cannot created table")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": table,
	})
}

func (t *TableController) TableGetAll(ctx *fiber.Ctx) error {
	tables, err := t.tableService.TableGetAll()
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failde", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": tables,
	})
}

func (t *TableController) TableGetById(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	table, err := t.tableService.TableGetById(uint(Id))
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": table,
	})
}

func (t *TableController) TableUpdate(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestTableUpdate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	table, err := t.tableService.TableUpdate(ctx, uint(Id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": table,
	})
}

func (t *TableController) TableDelete(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	if err := t.tableService.TableDelete(uint(Id)); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Table Successfully",
	})

}
