package controllers

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	service "api/the_deck/Service"
	utils "api/the_deck/Utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RequestTableController struct {
	requestTableService service.RequestTableService
}

func NewRequestTableController(rts service.RequestTableService) *RequestTableController {
	return &RequestTableController{requestTableService: rts}
}

func (rtc *RequestTableController) CreateRequestTable(ctx *fiber.Ctx) error {
	var input dto.AddRequestTable
	if err := ctx.BodyParser(&input); err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	requestTable, err := rtc.requestTableService.CreateRequest(ctx, input)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": requestTable,
	})
}

func (rtc *RequestTableController) GetMyReqTable(ctx *fiber.Ctx) error {
	customer := ctx.Locals("customer").(entity.Customer)
	req, err := rtc.requestTableService.GetMyReqTable(customer.Id)
	if err != nil {
		return utils.MessageJSON(ctx, 404, "Failed", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": req,
	})
}

func (rtc *RequestTableController) GetAllRequest(ctx *fiber.Ctx) error {
	req_table, err := rtc.requestTableService.GetAllRequestTable()
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": req_table,
	})
}

func (rtc *RequestTableController) UpdateStatus(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.MessageJSON(ctx, 400, "Failed", err.Error())
	}

	var input dto.UpdateRequestTable
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	req_table, err := rtc.requestTableService.UpdateStatus(uint(id), input)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": req_table,
	})
}
