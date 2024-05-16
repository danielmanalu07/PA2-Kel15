package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
	"errors"
	"os"
	"path/filepath"
)

type ProductRepository interface {
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]entity.Product, error)
	GetById(id uint) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id uint) error
}

type productRepository struct{}

const PathImageProduct = "./Public/Product"

func (p *productRepository) Delete(id uint) error {
	var product entity.Product

	if err := database.DB.First(&product, "id = ?", id).Error; err != nil {
		return errors.New("record not found")
	}

	if product.Image != "" {
		imagePath := filepath.Join(PathImageProduct, product.Image)
		if err := os.Remove(imagePath); err != nil {
			return err
		}
	}

	// Delete the product from the database
	if err := database.DB.Delete(&product, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
func (p *productRepository) Update(product *entity.Product) (*entity.Product, error) {
	if err := database.DB.Debug().Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) GetById(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := database.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) GetAll() ([]entity.Product, error) {
	var product []entity.Product
	if err := database.DB.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) Create(product entity.Product) (*entity.Product, error) {
	if err := database.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
