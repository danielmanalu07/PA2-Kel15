package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*response.AdminResponse, error)
	GetProfile(ctx *fiber.Ctx) (*response.AdminResponse, error)
	LogoutAdmin(ctx *fiber.Ctx) (*fiber.Cookie, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(ar repository.AdminRepository) AdminService {
	return &adminService{adminRepository: ar}
}

func (a *adminService) AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*response.AdminResponse, error) {
	admin, err := a.adminRepository.AdminLogin(ctx, input)
	if err != nil {
		return nil, err
	}

	claims := jwt.StandardClaims{
		Issuer:    admin.Username,
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Subject:   "admin",
	}

	tokens, err := utils.GenerateToken(&claims)
	if err != nil {
		return nil, err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokens,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	adminResponse := &response.AdminResponse{
		Username: admin.Username,
		Password: admin.Password,
	}

	return adminResponse, nil
}

func (a *adminService) GetProfile(ctx *fiber.Ctx) (*response.AdminResponse, error) {
	data, err := a.adminRepository.GetProfile(ctx)
	if err != nil {
		return nil, utils.MessageJSON(ctx, 401, "Failed", "Unauthenticated")
	}

	admins := &response.AdminResponse{
		Username: data.Username,
		Password: data.Password,
	}
	return admins, nil
}

func (a *adminService) LogoutAdmin(ctx *fiber.Ctx) (*fiber.Cookie, error) {
	cookie, err := a.adminRepository.LogoutAdmin(ctx)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}
