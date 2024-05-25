package repository

import (
	"api/the_deck/Models/entity"
	"errors"

	"gorm.io/gorm"
)

type RequestTableRepository interface {
	Create(requestTable entity.RequestTable) (*entity.RequestTable, error)
	GetAll() ([]entity.RequestTable, error)
	GetById(id uint) (*entity.RequestTable, error)
	Update(requestTable *entity.RequestTable) (*entity.RequestTable, error)
	Delete(id uint) error
}

type requestTableRepository struct {
	db *gorm.DB
}

func NewRequestTableRepository() RequestTableRepository {
	return &requestTableRepository{}
}

func (rtr *requestTableRepository) Create(requestTable entity.RequestTable) (*entity.RequestTable, error) {
	if err := rtr.db.Create(&requestTable).Error; err != nil {
		return nil, err
	}
	return &requestTable, nil
}

func (rtr *requestTableRepository) GetAll() ([]entity.RequestTable, error) {
	var requestTables []entity.RequestTable
	if err := rtr.db.Find(&requestTables).Error; err != nil {
		return nil, err
	}
	return requestTables, nil
}

func (rtr *requestTableRepository) GetById(id uint) (*entity.RequestTable, error) {
	var requestTable entity.RequestTable
	if err := rtr.db.First(&requestTable, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &requestTable, nil
}

func (rtr *requestTableRepository) Update(requestTable *entity.RequestTable) (*entity.RequestTable, error) {
	if err := rtr.db.Save(requestTable).Error; err != nil {
		return nil, err
	}
	return requestTable, nil
}

func (rtr *requestTableRepository) Delete(id uint) error {
	if err := rtr.db.Delete(&entity.RequestTable{}, id).Error; err != nil {
		return err
	}
	return nil
}
