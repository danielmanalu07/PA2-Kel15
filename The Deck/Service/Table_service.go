package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"

	"github.com/gofiber/fiber/v2"
)

type TableService interface {
	TableCreate(ctx *fiber.Ctx, input dto.RequestTableCreate) (*response.TableResponse, error)
	TableGetAll() ([]response.TableResponse, error)
	TableGetById(id uint) (*response.TableResponse, error)
	TableUpdate(ctx *fiber.Ctx, id uint, input dto.RequestTableUpdate) (*response.TableResponse, error)
	TableDelete(id uint) error
}

type tableService struct {
	tableService repository.TableRepository
}

func (t *tableService) TableDelete(id uint) error {
	return t.tableService.Delete(id)
}

func (t *tableService) TableUpdate(ctx *fiber.Ctx, id uint, input dto.RequestTableUpdate) (*response.TableResponse, error) {
	table, err := t.tableService.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Number != 0 {
		table.Number = input.Number
	}

	if input.Capacity != 0 {
		table.Capacity = input.Capacity
	}

	updateTbl, err := t.tableService.Update(table)
	if err != nil {
		return nil, err
	}

	respon := &response.TableResponse{
		Id:       updateTbl.Id,
		Number:   updateTbl.Number,
		Capacity: updateTbl.Capacity,
		AdminID:  updateTbl.AdminID,
		Status:   updateTbl.Status,
	}

	return respon, nil
}

func (t *tableService) TableGetById(id uint) (*response.TableResponse, error) {
	table, err := t.tableService.GetById(id)
	if err != nil {
		return nil, err
	}

	tbl := &response.TableResponse{
		Id:       table.Id,
		Number:   table.Number,
		Capacity: table.Capacity,
		AdminID:  table.AdminID,
		Status:   table.Status,
	}

	return tbl, nil
}

func (t *tableService) TableGetAll() ([]response.TableResponse, error) {
	tables, err := t.tableService.GetAll()
	if err != nil {
		return nil, err
	}

	var respon []response.TableResponse
	for _, table := range tables {
		respon = append(respon, response.TableResponse{
			Id:       table.Id,
			Number:   table.Number,
			Capacity: table.Capacity,
			AdminID:  table.AdminID,
			Status:   table.Status,
		})
	}
	return respon, nil
}

func (t *tableService) TableCreate(ctx *fiber.Ctx, input dto.RequestTableCreate) (*response.TableResponse, error) {
	admin := ctx.Locals("admin").(entity.Admin)
	table := entity.Table{
		Number:   input.Number,
		Capacity: input.Capacity,
		AdminID:  admin.Id,
		Admin:    admin,
		Status:   0,
	}

	tableCreate, err := t.tableService.Create(table)
	if err != nil {
		return nil, err
	}

	respon := &response.TableResponse{
		Id:       tableCreate.Id,
		Number:   tableCreate.Number,
		Capacity: tableCreate.Capacity,
		AdminID:  tableCreate.AdminID,
		Status:   tableCreate.Status,
	}

	return respon, nil
}

func NewTableService(tr repository.TableRepository) TableService {
	return &tableService{tableService: tr}
}
