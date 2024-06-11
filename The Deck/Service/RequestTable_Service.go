package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"

	"github.com/gofiber/fiber/v2"
)

type RequestTableService interface {
	CreateRequest(ctx *fiber.Ctx, input dto.AddRequestTable) (*response.ResponseRequestTable, error)
	GetMyReqTable(customerId uint) ([]response.ResponseRequestTable, error)
	GetAllRequestTable() ([]response.ResponseRequestTable, error)
	UpdateStatus(id uint, input dto.UpdateRequestTable) (*response.ResponseRequestTable, error)
}

type requestTableService struct {
	requestTableRepository repository.RequestTableRepository
}

func (r *requestTableService) UpdateStatus(id uint, input dto.UpdateRequestTable) (*response.ResponseRequestTable, error) {
	req_table, err := r.requestTableRepository.GetReqById(id)
	if err != nil {
		return nil, err
	}

	if input.Status != 0 {
		req_table.Status = input.Status
	}

	save, err := r.requestTableRepository.Update(req_table)
	if err != nil {
		return nil, err
	}

	response := &response.ResponseRequestTable{
		Id:         save.Id,
		TableID:    save.TableID,
		CustomerID: save.CustomerID,
		AdminID:    save.AdminID,
		Notes:      save.Notes,
		Status:     save.Status,
	}

	return response, nil
}

func (r *requestTableService) GetAllRequestTable() ([]response.ResponseRequestTable, error) {
	req_table, err := r.requestTableRepository.GetAllRequest()
	if err != nil {
		return nil, err
	}

	var req_tableResponses []response.ResponseRequestTable
	for _, req_tables := range req_table {
		req_tableResponse := response.ResponseRequestTable{
			Id:         req_tables.Id,
			CustomerID: req_tables.CustomerID,
			TableID:    req_tables.TableID,
			AdminID:    req_tables.AdminID,
			Status:     req_tables.Status,
			Notes:      req_tables.Notes,
		}

		req_tableResponses = append(req_tableResponses, req_tableResponse)
	}

	return req_tableResponses, nil
}

func (r *requestTableService) GetMyReqTable(customerId uint) ([]response.ResponseRequestTable, error) {
	req_table, err := r.requestTableRepository.GetMyReq(customerId)
	if err != nil {
		return nil, err
	}

	var req_tableResponses []response.ResponseRequestTable
	for _, req_tables := range req_table {
		req_tableResponse := response.ResponseRequestTable{
			Id:         req_tables.Id,
			CustomerID: req_tables.CustomerID,
			TableID:    req_tables.TableID,
			AdminID:    req_tables.AdminID,
			Status:     req_tables.Status,
			Notes:      req_tables.Notes,
		}

		req_tableResponses = append(req_tableResponses, req_tableResponse)
	}

	return req_tableResponses, nil
}

func (r *requestTableService) CreateRequest(ctx *fiber.Ctx, input dto.AddRequestTable) (*response.ResponseRequestTable, error) {
	customer, ok := ctx.Locals("customer").(entity.Customer)
	if !ok {
		return nil, utils.MessageJSON(ctx, 400, "error", "Customer not found")
	}

	rt := entity.RequestTable{
		TableID:    input.TableId,
		CustomerID: customer.Id,
		Notes:      input.Notes,
		AdminID:    nil,
		Status:     0,
	}

	save, err := r.requestTableRepository.Create(rt)
	if err != nil {
		return nil, err
	}

	rtResponse := &response.ResponseRequestTable{
		Id:         save.Id,
		CustomerID: save.CustomerID,
		TableID:    save.TableID,
		Status:     save.Status,
		Notes:      save.Notes,
	}

	return rtResponse, nil
}

func NewRequestTableService(rt repository.RequestTableRepository) RequestTableService {
	return &requestTableService{requestTableRepository: rt}
}
