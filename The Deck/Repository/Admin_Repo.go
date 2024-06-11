package repository

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	utils "api/the_deck/Utils"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminRepository interface {
	AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*entity.Admin, error)
	GetProfile(ctx *fiber.Ctx) (*entity.Admin, error)
	LogoutAdmin(ctx *fiber.Ctx) (*fiber.Cookie, error)
	Approve(order entity.Order) (*entity.Order, error)
	ApproveTable(tx *gorm.DB, reqTable entity.RequestTable) (*entity.RequestTable, error)
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (a *adminRepository) ApproveTable(tx *gorm.DB, reqTable entity.RequestTable) (*entity.RequestTable, error) {
	if err := tx.Save(&reqTable).Error; err != nil {
		return nil, err
	}

	return &reqTable, nil
}

func (a *adminRepository) Approve(order entity.Order) (*entity.Order, error) {
	if err := database.DB.Save(order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
func (a *adminRepository) AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*entity.Admin, error) {
	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return nil, err
	}
	var admin entity.Admin
	result := database.DB.First(&admin, "username = ?", input.Username)
	if result.Error != nil {
		return nil, result.Error
	}

	// Here you should check the password using your utility function
	if !utils.CheckPassword(input.Password, admin.Password) {
		return nil, errors.New("incorrect password")
	}

	return &admin, nil
}

func (a *adminRepository) GetProfile(ctx *fiber.Ctx) (*entity.Admin, error) {
	admin := ctx.Locals("admin").(entity.Admin)
	return &admin, nil
}

func (a *adminRepository) LogoutAdmin(ctx *fiber.Ctx) (*fiber.Cookie, error) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	return &cookie, nil
}
