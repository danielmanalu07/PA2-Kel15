package service

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*response.AdminResponse, string, error)
	GetProfile(ctx *fiber.Ctx) (*response.AdminResponse, error)
	LogoutAdmin(ctx *fiber.Ctx) (*fiber.Cookie, error)
	UpdateStatus(ctx *fiber.Ctx, id uint, input dto.RequestOrderUpdateStatus) (*response.OrderResponse, error)
	ApproveReqTable(ctx *fiber.Ctx, id uint, input dto.UpdateRequestTable) (*response.ResponseRequestTable, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(ar repository.AdminRepository) AdminService {
	return &adminService{adminRepository: ar}
}

func (a *adminService) ApproveReqTable(ctx *fiber.Ctx, id uint, input dto.UpdateRequestTable) (*response.ResponseRequestTable, error) {
	// var req_table entity.RequestTable
	// if err := database.DB.First(&req_table, "id = ?", id).Error; err != nil {
	// 	return nil, err
	// }

	// admin := ctx.Locals("admin").(entity.Admin)

	// if input.Status != 0 {
	// 	req_table.Status = input.Status
	// }

	// req_table.AdminID = &admin.Id

	// save, err := a.adminRepository.ApproveTable(req_table)
	// if err != nil {
	// 	return nil, err
	// }

	// response := response.ResponseRequestTable{
	// 	Id:         save.Id,
	// 	CustomerID: save.CustomerID,
	// 	TableID:    save.TableID,
	// 	AdminID:    save.AdminID,
	// 	Notes:      save.Notes,
	// 	Status:     save.Status,
	// }

	// return &response, nil
	var reqTable entity.RequestTable
	if err := database.DB.First(&reqTable, "id = ?", id).Error; err != nil {
		return nil, err
	}

	admin := ctx.Locals("admin").(entity.Admin)

	if input.Status != 0 {
		reqTable.Status = input.Status
	}

	reqTable.AdminID = &admin.Id

	// Mulai transaksi
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	save, err := a.adminRepository.ApproveTable(tx, reqTable)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if input.Status == 1 {
		var table entity.Table
		if err := tx.First(&table, "id = ?", reqTable.TableID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		table.Status = 1
		if err := tx.Save(&table).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if input.Status == 3 {
		var table entity.Table
		if err := tx.First(&table, "id = ?", reqTable.TableID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		table.Status = 0
		if err := tx.Save(&table).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	response := response.ResponseRequestTable{
		Id:         save.Id,
		CustomerID: save.CustomerID,
		TableID:    save.TableID,
		AdminID:    save.AdminID,
		Notes:      save.Notes,
		Status:     save.Status,
	}

	return &response, nil
}

func (a *adminService) UpdateStatus(ctx *fiber.Ctx, id uint, input dto.RequestOrderUpdateStatus) (*response.OrderResponse, error) {
	var order entity.Order
	if err := database.DB.First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	admin := ctx.Locals("admin").(entity.Admin)

	if input.Status != 0 {
		order.Status = input.Status
	}

	order.AdminID = &admin.Id

	save, err := a.adminRepository.Approve(order)
	if err != nil {
		return nil, err
	}

	response := response.OrderResponse{
		Id:             save.Id,
		Code:           save.Code,
		CustomerID:     save.CustomerID,
		Products:       save.Products,
		Total:          save.Total,
		Note:           save.Note,
		PaymentMethod:  save.PaymentMethod,
		PickUpType:     save.PickUpType,
		ProofOfPayment: save.ProofOfPayment,
		Status:         save.Status,
	}

	return &response, nil
}

func (a *adminService) AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*response.AdminResponse, string, error) {
	admin, err := a.adminRepository.AdminLogin(ctx, input)
	if err != nil {
		return nil, "", err
	}

	claims := jwt.StandardClaims{
		Issuer:    admin.Username,
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Subject:   "admin",
	}

	tokens, err := utils.GenerateToken(&claims)
	if err != nil {
		return nil, "", err
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

	return adminResponse, tokens, nil
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
