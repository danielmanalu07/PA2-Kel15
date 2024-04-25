package implements

import (
	"os"
	"pa2/database"
	"pa2/models/dto"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/service"
	"pa2/utils"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const PathImageProduct = "./Public/file_product"

type ProductImpl struct{}

// DeleteProduct implements service.ProductService.
func (c *ProductImpl) DeleteProduct(ctx *fiber.Ctx, product *entity.Product, id int) error {
	err := database.DB.Debug().First(&product, "id = ? ", id).Error
	if err != nil {
		return response.BadRequestResponse(ctx)
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return response.InternalServerError(ctx)
	}

	if product.Image != "" {
		imagePath := filepath.Join(PathImageProduct, product.Image)
		if err := os.Remove(imagePath); err != nil {
			return response.BadRequestResponse(ctx)
		}
	}

	return err
}

func (c *ProductImpl) UpdateProduct(ctx *fiber.Ctx, product *entity.Product, id int) error {
	inputReq := new(dto.RequestUpdateProduct)
	if err := ctx.BodyParser(inputReq); err != nil {
		return err
	}

	// Mendapatkan data produk berdasarkan ID
	result := database.DB.First(&product, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return response.NotFoundResponse(ctx)
	}

	if inputReq.Name != "" {
		product.Name = inputReq.Name
	}
	if inputReq.Description != "" {
		product.Description = inputReq.Description
	}
	if inputReq.Price != 0 {
		product.Price = inputReq.Price
	}
	if inputReq.CategoryId != 0 {
		product.CategoryId = inputReq.CategoryId
	}

	// Memeriksa apakah file gambar baru diunggah
	newImage, err := ctx.FormFile("image")
	if err == nil {
		// Menghapus file gambar lama jika ada
		if product.Image != "" {
			oldImagePath := filepath.Join(PathImageProduct, product.Image)
			os.Remove(oldImagePath)
		}

		// Menyimpan file gambar baru
		newFileName := utils.GenerateImageFile(product.Name, newImage.Filename)
		if err := ctx.SaveFile(newImage, filepath.Join(PathImageProduct, newFileName)); err != nil {
			return err
		}

		product.Image = newFileName
	}

	// Memperbarui data produk di database
	if err := database.DB.Save(&product).Error; err != nil {
		return response.InternalServerError(ctx)
	}

	return nil
}

func (c *ProductImpl) GetProductById(product *entity.Product, id int) error {
	result := database.DB.First(&product, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (c *ProductImpl) GetAllProduct(product *[]entity.Product) error {
	result := database.DB.Find(&product)
	if result == nil {
		return response.NotFoundResponse(&fiber.Ctx{})
	}

	return result.Error
}

func (c *ProductImpl) CreateProduct(ctx *fiber.Ctx) error {
	inputReq := new(dto.RequestCreateProduct)
	if err := ctx.BodyParser(inputReq); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(inputReq); err != nil {
		return response.BadRequestResponse(ctx)
	}

	CategoryId, err := strconv.Atoi(ctx.FormValue("category_id"))
	if err != nil {
		return response.BadRequestResponse(ctx)
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return response.BadRequestResponse(ctx)
	}

	filename := utils.GenerateImageFile(inputReq.Name, image.Filename)

	if err := ctx.SaveFile(image, filepath.Join(PathImageProduct, filename)); err != nil {
		return response.InternalServerError(ctx)
	}

	admin := ctx.Locals("admin").(entity.Admin)

	products := &entity.Product{
		Name:        inputReq.Name,
		Description: inputReq.Description,
		Price:       inputReq.Price,
		Image:       filename,
		CategoryId:  uint(CategoryId),
		AdminID:     admin.Id,
	}

	result := database.DB.Create(products)
	return result.Error
}

func NewProductRepository() service.ProductService {
	return &ProductImpl{}
}
