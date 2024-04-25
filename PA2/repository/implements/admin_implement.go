package implements

import (
	"context"
	"database/sql"
	"errors"
	"pa2/database"
	"pa2/models/dto"
	"pa2/models/entity"
	"pa2/repository/service"
	"pa2/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AdminImpl struct{}

func (a *AdminImpl) AdminLogout(ctx *fiber.Ctx) (string, error) {
	cookie := fiber.Cookie{
		Name:     "admin",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	ctx.Locals("admin", nil)

	return "", nil
}

func (a *AdminImpl) AdminGetProfile(ctx *fiber.Ctx) (*entity.Admin, error) {
	admin := ctx.Locals("admin").(entity.Admin)
	return &admin, nil
}

func (a *AdminImpl) AdminLogin(ctx context.Context, credentials dto.RequestAdminLogin) (string, error) {
	if err := validator.New().Struct(credentials); err != nil {
		return "", errors.New("invalid input")
	}

	var admin entity.Admin
	err := database.DB.ConnPool.QueryRowContext(ctx, "SELECT id, username, password FROM admins WHERE username = ?", credentials.Username).Scan(&admin.Id, &admin.Username, &admin.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	if !utils.CheckHashPassword(credentials.Password, admin.Password) {
		return "", errors.New("invalid username or password")
	}

	claims := jwt.StandardClaims{
		Issuer:    admin.Username,
		Subject:   strconv.Itoa(int(admin.Id)),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AdminImpl) AdminCreate(ctx context.Context, admin *entity.Admin) error {
	_, err := database.DB.ConnPool.ExecContext(ctx, "INSERT INTO admins (username, password, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)", admin.Username, admin.Password)
	return err
}

func NewAdminRepository() service.AdminService {
	return &AdminImpl{}
}
