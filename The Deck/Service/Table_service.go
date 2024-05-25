package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"

	"github.com/gofiber/fiber/v2"
)

type TableService interface {
	TableCreate(input dto.RequestTableCreate) (*response.TableResponse, error)
	TableGetAll() ([]response.TableResponse, error)
	TableGetById(id uint) (*response.TableResponse, error)
	TableUpdate(ctx *fiber.Ctx, id uint, input dto.RequestTableUpdate) (*response.TableResponse, error)
	TableDelete(id uint) error
}

type tableService struct {
	tableRepository repository.TableRepository
}

func (t *tableService) TableDelete(id uint) error {
	return t.tableRepository.Delete(id)
}

func (t *tableService) TableUpdate(ctx *fiber.Ctx, id uint, input dto.RequestTableUpdate) (*response.TableResponse, error) {
	table, err := t.tableRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Status != "" {
		table.Status = input.Status
	}

	updatedTable, err := t.tableRepository.Update(table)
	if err != nil {
		return nil, err
	}

	respon := &response.TableResponse{
		Id:       updatedTable.Id,
		Number:   updatedTable.Number,
		Capacity: updatedTable.Capacity,
		Status:   updatedTable.Status,
	}

	return respon, nil
}

func (t *tableService) TableGetById(id uint) (*response.TableResponse, error) {
	table, err := t.tableRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	tbl := &response.TableResponse{
		Id:       table.Id,
		Number:   table.Number,
		Capacity: table.Capacity,
		Status:   table.Status,
	}

	return tbl, nil
}

func (t *tableService) TableGetAll() ([]response.TableResponse, error) {
	tables, err := t.tableRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var respon []response.TableResponse
	for _, table := range tables {
		respon = append(respon, response.TableResponse{
			Id:       table.Id,
			Number:   table.Number,
			Capacity: table.Capacity,
			Status:   table.Status,
		})
	}
	return respon, nil
}

func (t *tableService) TableCreate(input dto.RequestTableCreate) (*response.TableResponse, error) {
	// Set the default status to "pending" if not provided
	status := input.Status
	if status == "" {
		status = "kosong"
	}

	table := entity.Table{
		Status:   status,
	}

	createdTable, err := t.tableRepository.Create(table)
	if err != nil {
		return nil, err
	}

	respon := &response.TableResponse{
		Id:       createdTable.Id,
		Number:   createdTable.Number,
		Capacity: createdTable.Capacity,
		Status:   createdTable.Status,
	}

	return respon, nil
}

func NewTableService(tr repository.TableRepository) TableService {
	return &tableService{tableRepository: tr}
}
