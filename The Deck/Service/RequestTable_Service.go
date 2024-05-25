package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RequestTableService interface {
	CreateRequestTable(ctx *fiber.Ctx, input dto.RequestTableCreate, customerID uint) (*response.RequestTableResponse, error)
	GetAllRequestTables() ([]response.RequestTableResponse, error)
	GetRequestTableById(id uint) (*response.RequestTableResponse, error)
	UpdateRequestTable(id uint, input dto.RequestTableUpdate) (*response.RequestTableResponse, error)
	DeleteRequestTable(id uint) error
}

type requestTableService struct {
	requestTableRepository repository.RequestTableRepository
}

func NewRequestTableService(rqt repository.RequestTableRepository) RequestTableService {
	return &requestTableService{requestTableRepository: rqt}
}

func (rts *requestTableService) DeleteRequestTable(id uint) error {
	return rts.requestTableRepository.Delete(id)
}

func (rts *requestTableService) CreateRequestTable(ctx *fiber.Ctx, input dto.RequestTableCreate, customerID uint) (*response.RequestTableResponse, error) {
	Status := input.Status
	if Status == "" {
		Status = "pending"
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return nil, err
	}
	requestTable := entity.RequestTable{
		Description: input.Description,
		Status:      Status,
		StartDate:   startDate,
		EndDate:     endDate,
		TableID:     input.TableID,
		CustomerID:  customerID,
	}

	createdRequestTable, err := rts.requestTableRepository.Create(requestTable)
	if err != nil {
		return nil, err
	}

	respon := &response.RequestTableResponse{
		Id:          createdRequestTable.Id,
		Description: createdRequestTable.Description,
		Status:      createdRequestTable.Status,
		StartDate:   createdRequestTable.StartDate,
		EndDate:     createdRequestTable.EndDate,
		TableID:     createdRequestTable.TableID,
		CustomerID:  createdRequestTable.CustomerID,
		CreatedAt:   createdRequestTable.CreatedAt,
		UpdatedAt:   createdRequestTable.UpdatedAt,
	}

	return respon, nil
}

func (rts *requestTableService) GetAllRequestTables() ([]response.RequestTableResponse, error) {
	requestTables, err := rts.requestTableRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var respon []response.RequestTableResponse
	for _, requestTable := range requestTables {
		respon = append(respon, response.RequestTableResponse{
			Id:          requestTable.Id,
			Description: requestTable.Description,
			Status:      requestTable.Status,
			StartDate:   requestTable.StartDate,
			EndDate:     requestTable.EndDate,
			TableID:     requestTable.TableID,
			CustomerID:  requestTable.CustomerID,
			CreatedAt:   requestTable.CreatedAt,
			UpdatedAt:   requestTable.UpdatedAt,
		})
	}
	return respon, nil
}

func (rts *requestTableService) GetRequestTableById(id uint) (*response.RequestTableResponse, error) {
	requestTable, err := rts.requestTableRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	respon := &response.RequestTableResponse{
		Id:          requestTable.Id,
		Description: requestTable.Description,
		Status:      requestTable.Status,
		StartDate:   requestTable.StartDate,
		EndDate:     requestTable.EndDate,
		TableID:     requestTable.TableID,
		CustomerID:  requestTable.CustomerID,
		CreatedAt:   requestTable.CreatedAt,
		UpdatedAt:   requestTable.UpdatedAt,
	}

	return respon, nil
}

func (rts *requestTableService) UpdateRequestTable(id uint, input dto.RequestTableUpdate) (*response.RequestTableResponse, error) {
	requestTable, err := rts.requestTableRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Description != "" {
		requestTable.Description = input.Description
	}

	if input.Status != "" {
		requestTable.Status = input.Status
	}

	if input.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", input.StartDate)
		if err != nil {
			return nil, err
		}
		requestTable.StartDate = startDate
	}

	if input.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			return nil, err
		}
		requestTable.EndDate = endDate
	}

	if input.TableID != 0 {
		requestTable.TableID = input.TableID
	}

	updatedRequestTable, err := rts.requestTableRepository.Update(requestTable)
	if err != nil {
		return nil, err
	}

	respon := &response.RequestTableResponse{
		Id:          updatedRequestTable.Id,
		Description: updatedRequestTable.Description,
		Status:      updatedRequestTable.Status,
		StartDate:   updatedRequestTable.StartDate,
		EndDate:     updatedRequestTable.EndDate,
		TableID:     updatedRequestTable.TableID,
		CustomerID:  updatedRequestTable.CustomerID,
		CreatedAt:   updatedRequestTable.CreatedAt,
	}

	return respon, nil
}
