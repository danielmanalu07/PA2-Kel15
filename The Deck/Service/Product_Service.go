package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const PathImageProduct = "./Public/Product"

type ProductService interface {
	ProductCreate(ctx *fiber.Ctx, input dto.RequestProductCreate) (*response.ProductResponse, error)
	ProductGetAll() ([]response.ProductResponse, error)
	ProductGetById(id uint) (*response.ProductResponse, error)
	ProductUpdate(ctx *fiber.Ctx, id uint, input dto.RequestProductUpdate) (*response.ProductResponse, error)
	ProductDelete(id uint) error
	ProductGetByCategory(categoryId uint) ([]response.ProductResponse, error)
}

type productService struct {
	productService repository.ProductRepository
}

func (p *productService) ProductGetByCategory(categoryId uint) ([]response.ProductResponse, error) {
	products, err := p.productService.GetByCategory(categoryId)
	if err != nil {
		return nil, err
	}

	var responses []response.ProductResponse
	for _, product := range products {
		response := response.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
			CategoryID:  product.CategoryID,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *productService) ProductDelete(id uint) error {
	err := p.productService.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("record not found")
		}
		return err
	}
	return nil
}

func (p *productService) ProductUpdate(ctx *fiber.Ctx, id uint, input dto.RequestProductUpdate) (*response.ProductResponse, error) {
	product, err := p.productService.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		product.Name = input.Name
	}

	if input.Description != "" {
		product.Description = input.Description
	}

	if input.Price != "" {
		product.Price = input.Price
	}

	if input.CategoryID != 0 {
		product.CategoryID = input.CategoryID
	}

	newImage, err := ctx.FormFile("image")
	if err == nil {
		if product.Image != "" {
			oldPath := filepath.Join(PathImageProduct, product.Image)
			os.Remove(oldPath)
		}

		newFilename := utils.GenerateImageFile(product.Name, newImage.Filename)
		if err := ctx.SaveFile(newImage, filepath.Join(PathImageProduct, newFilename)); err != nil {
			return nil, err
		}

		product.Image = newFilename
	}

	updateProduct, err := p.productService.Update(product)
	if err != nil {
		return nil, err
	}

	respon := &response.ProductResponse{
		Id:          updateProduct.Id,
		Name:        updateProduct.Name,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
		Image:       updateProduct.Image,
		CategoryID:  updateProduct.CategoryID,
		AdminID:     updateProduct.AdminID,
	}

	return respon, nil
}

func (p *productService) ProductGetById(id uint) (*response.ProductResponse, error) {
	prod, err := p.productService.GetById(id)
	if err != nil {
		return nil, err
	}

	product := &response.ProductResponse{
		Id:          prod.Id,
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
		Image:       prod.Image,
		CategoryID:  prod.CategoryID,
		AdminID:     prod.AdminID,
	}

	return product, nil
}

func (p *productService) ProductGetAll() ([]response.ProductResponse, error) {
	product, err := p.productService.GetAll()
	if err != nil {
		return nil, err
	}

	var respon []response.ProductResponse
	for _, products := range product {
		respon = append(respon, response.ProductResponse{
			Id:          products.Id,
			Name:        products.Name,
			Description: products.Description,
			Price:       products.Price,
			Image:       products.Image,
			CategoryID:  products.CategoryID,
			AdminID:     products.AdminID,
		})
	}

	return respon, nil
}

func (p *productService) ProductCreate(ctx *fiber.Ctx, input dto.RequestProductCreate) (*response.ProductResponse, error) {
	image, err := ctx.FormFile("image")
	if err != nil {
		return nil, err
	}

	filename := utils.GenerateImageFile(input.Name, image.Filename)

	if err := ctx.SaveFile(image, filepath.Join(PathImageProduct, filename)); err != nil {
		return nil, err
	}

	CategoryId, err := strconv.Atoi(ctx.FormValue("category_id"))
	if err != nil {
		return nil, err
	}

	input.CategoryID = uint(CategoryId)

	admin := ctx.Locals("admin").(entity.Admin)

	product := entity.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Image:       filename,
		CategoryID:  uint(CategoryId),
		AdminID:     admin.Id,
		Admin:       admin,
	}

	createProduct, err := p.productService.Create(product)
	if err != nil {
		return nil, err
	}

	respon := &response.ProductResponse{
		Id:          createProduct.Id,
		Name:        createProduct.Name,
		Description: createProduct.Description,
		Price:       createProduct.Price,
		Image:       createProduct.Image,
		CategoryID:  createProduct.CategoryID,
		AdminID:     createProduct.AdminID,
	}

	return respon, nil
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{productService: pr}
}
