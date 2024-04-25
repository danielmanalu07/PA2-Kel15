package service

import (
	"context"
	"pa2/models/dto"
	"pa2/models/entity"

	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	AdminCreate(ctx context.Context, admin *entity.Admin) error
	AdminLogin(ctx context.Context, credentials dto.RequestAdminLogin) (string, error)
	AdminGetProfile(ctx *fiber.Ctx) (*entity.Admin, error)
	AdminLogout(ctx *fiber.Ctx) (string, error)
}
