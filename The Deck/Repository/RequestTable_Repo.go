package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type RequestTableRepository interface {
	Create(requestTable entity.RequestTable) (*entity.RequestTable, error)
	GetAllRequest() ([]entity.RequestTable, error)
	GetMyReq(customerId uint) ([]entity.RequestTable, error)
	GetReqById(id uint) (*entity.RequestTable, error)
	Update(reqTable *entity.RequestTable) (*entity.RequestTable, error)
}

type requestTableRepository struct{}

func (r *requestTableRepository) Update(reqTable *entity.RequestTable) (*entity.RequestTable, error) {
	if err := database.DB.Save(reqTable).Error; err != nil {
		return nil, err
	}

	return reqTable, nil
}

func (r *requestTableRepository) GetReqById(id uint) (*entity.RequestTable, error) {
	var req_table entity.RequestTable
	if err := database.DB.First(&req_table, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &req_table, nil
}

func (r *requestTableRepository) GetAllRequest() ([]entity.RequestTable, error) {
	var req_table []entity.RequestTable
	if err := database.DB.Find(&req_table).Error; err != nil {
		return nil, err
	}

	return req_table, nil
}

func (r *requestTableRepository) GetMyReq(customerId uint) ([]entity.RequestTable, error) {
	var req_table []entity.RequestTable
	if err := database.DB.Where("customer_id = ?", customerId).Find(&req_table).Error; err != nil {
		return nil, err
	}

	return req_table, nil
}

func (r *requestTableRepository) Create(requestTable entity.RequestTable) (*entity.RequestTable, error) {
	if err := database.DB.Debug().Create(&requestTable).Error; err != nil {
		return nil, err
	}

	return &requestTable, nil
}

func NewRequestTableRepository() RequestTableRepository {
	return &requestTableRepository{}
}
