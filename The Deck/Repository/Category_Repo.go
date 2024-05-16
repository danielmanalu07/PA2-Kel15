package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
)

type CategoryRepository interface {
	Create(category entity.Category) (*entity.Category, error)
	GetAll() ([]entity.Category, error)
	GetById(id uint) (*entity.Category, error)
	Update(category *entity.Category) (*entity.Category, error)
	Delete(id uint) error
}

type categoryRepository struct{}

func (c *categoryRepository) Delete(id uint) error {
	var category entity.Category
	if err := database.DB.Debug().Delete(&category, "id = ?", id); err != nil {
		return err.Error
	}
	return nil
}

func (c *categoryRepository) Update(category *entity.Category) (*entity.Category, error) {
	if err := database.DB.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) GetById(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categoryRepository) GetAll() ([]entity.Category, error) {
	var category []entity.Category
	if err := database.DB.Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) Create(category entity.Category) (*entity.Category, error) {
	if err := database.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}
