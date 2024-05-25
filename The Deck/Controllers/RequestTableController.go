package controllers

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Service"
	"api/the_deck/Utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RequestTableController struct {
	requestTableService service.RequestTableService
}

func NewRequestTableController(rts service.RequestTableService) *RequestTableController {
	return &RequestTableController{requestTableService: rts}
}

func (rtc *RequestTableController) CreateRequestTable(ctx *fiber.Ctx) error {
	var input dto.RequestTableCreate
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return err
	}

	customerID := ctx.Locals("customerID").(uint)

	requestTable, err := rtc.requestTableService.CreateRequestTable(ctx, input, customerID)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", "Cannot create request table")
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": requestTable,
	})
}

func (rtc *RequestTableController) GetAllRequestTables(ctx *fiber.Ctx) error {
	requestTables, err := rtc.requestTableService.GetAllRequestTables()
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": requestTables,
	})
}

func (rtc *RequestTableController) GetRequestTableById(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	requestTable, err := rtc.requestTableService.GetRequestTableById(uint(Id))
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": requestTable,
	})
}

func (rtc *RequestTableController) UpdateRequestTable(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.RequestTableUpdate
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	requestTable, err := rtc.requestTableService.UpdateRequestTable(uint(Id), input)
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": requestTable,
	})
}

func (rtc *RequestTableController) DeleteRequestTable(ctx *fiber.Ctx) error {
	Id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	if err := rtc.requestTableService.DeleteRequestTable(uint(Id)); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Request Table Successfully",
	})
}
