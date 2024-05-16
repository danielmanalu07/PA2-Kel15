package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type TableRepository interface {
	Create(table entity.Table) (*entity.Table, error)
	GetAll() ([]entity.Table, error)
	GetById(id uint) (*entity.Table, error)
	Update(table *entity.Table) (*entity.Table, error)
	Delete(id uint) error
}

type tableRepository struct{}

func (t *tableRepository) Delete(id uint) error {
	var table entity.Table
	if err := database.DB.Debug().Delete(&table, "id = ?", id); err != nil {
		return err.Error
	}
	return nil
}

func (t *tableRepository) Update(table *entity.Table) (*entity.Table, error) {
	if err := database.DB.Debug().Save(table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

func (t *tableRepository) GetById(id uint) (*entity.Table, error) {
	var table entity.Table
	if err := database.DB.First(&table, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &table, nil
}

func (t *tableRepository) GetAll() ([]entity.Table, error) {
	var table []entity.Table
	if err := database.DB.Find(&table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

func (t *tableRepository) Create(table entity.Table) (*entity.Table, error) {
	if err := database.DB.Create(&table).Error; err != nil {
		return nil, err
	}

	return &table, nil
}

func NewTableRepository() TableRepository {
	return &tableRepository{}
}
