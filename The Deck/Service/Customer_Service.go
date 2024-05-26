package service

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"
	utils "api/the_deck/Utils"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const PathImageCustomer = "./Public/Customer"

type CustomerService interface {
	CustomerRegister(ctx *fiber.Ctx, input dto.RequestCustomerRegister) (*response.CustomerResponse, error)
	CustomerLogin(ctx *fiber.Ctx, input dto.RequestCustomerLogin) (*response.CustomerResponse, string, error)
	GetProfile(ctx *fiber.Ctx) (*response.CustomerResponse, error)
	CustomerLogout(ctx *fiber.Ctx) (*fiber.Cookie, error)
	CustomerUpdate(ctx *fiber.Ctx, input dto.RequestCustomerUpdateProfile) (*response.CustomerResponse, error)
	CustomerEditPassword(ctx *fiber.Ctx, input dto.RequestCustomerEditPassword) (*response.CustomerResponse, error)
	CustomerForgotPassword(ctx *fiber.Ctx, input dto.RequestCustomerForgotPassword) (*response.CustomerResponse, error)
}

type customerService struct {
	customerService repository.CustomerRepository
}

func (c *customerService) CustomerForgotPassword(ctx *fiber.Ctx, input dto.RequestCustomerForgotPassword) (*response.CustomerResponse, error) {
	customer := ctx.Locals("customer").(entity.Customer)
	if err := database.DB.Where("email = ?", input.Email).First(&customer).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Where("phone = ?", input.Phone).First(&customer).Error; err != nil {
		return nil, err
	}

	newPw, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return nil, err
	}

	if input.Password != "" {
		customer.Password = newPw
	}

	result, err := c.customerService.Forgot(&customer)
	if err != nil {
		return nil, err
	}

	response := &response.CustomerResponse{
		Id:          result.Id,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Password:    result.Password,
		Phone:       result.Phone,
		Address:     result.Address,
		Gender:      result.Gender,
		DateOfBirth: result.DateOfBirth,
		Image:       result.Image,
	}

	return response, nil
}

func (c *customerService) CustomerEditPassword(ctx *fiber.Ctx, input dto.RequestCustomerEditPassword) (*response.CustomerResponse, error) {
	customer := ctx.Locals("customer").(entity.Customer)
	if err := database.DB.Where("phone = ?", input.Phone).First(&customer).Error; err != nil {
		return nil, err
	}

	if err := database.DB.First(&customer, "id = ?", customer.Id).Error; err != nil {
		return nil, err
	}

	newPw, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return nil, err
	}

	if input.Password != "" {
		customer.Password = newPw
	}

	result, err := c.customerService.EditPW(&customer)
	if err != nil {
		return nil, err
	}

	response := &response.CustomerResponse{
		Id:          result.Id,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Password:    result.Password,
		Phone:       result.Phone,
		Address:     result.Address,
		Gender:      result.Gender,
		DateOfBirth: result.DateOfBirth,
		Image:       result.Image,
	}

	return response, nil

}

func (c *customerService) CustomerUpdate(ctx *fiber.Ctx, input dto.RequestCustomerUpdateProfile) (*response.CustomerResponse, error) {
	customer := ctx.Locals("customer").(entity.Customer)

	if input.Name != "" {
		customer.Name = input.Name
	}
	if input.Username != "" {
		customer.Username = input.Username
	}

	if input.Email != "" {
		customer.Email = input.Email
	}
	if input.Phone != "" {
		customer.Phone = input.Phone
	}
	if input.Address != "" {
		customer.Address = input.Address
	}
	if input.Gender != "" {
		customer.Gender = input.Gender
	}
	if input.DateOfBirth != "" {
		customer.DateOfBirth = input.DateOfBirth
	}

	newImage, err := ctx.FormFile("image")
	if err == nil {
		if customer.Image != "" {
			oldPath := filepath.Join(PathImageCustomer, customer.Image)
			os.Remove(oldPath)
		}

		newFilename := utils.GenerateImageFile(customer.Username, newImage.Filename)
		if err := ctx.SaveFile(newImage, filepath.Join(PathImageCustomer, newFilename)); err != nil {
			return nil, err
		}

		customer.Image = newFilename
	}

	save, err := c.customerService.Update(&customer)
	if err != nil {
		return nil, err
	}

	response := &response.CustomerResponse{
		Id:          save.Id,
		Name:        save.Name,
		Username:    save.Username,
		Email:       save.Email,
		Password:    save.Password,
		Phone:       save.Phone,
		Address:     save.Address,
		Gender:      save.Gender,
		DateOfBirth: save.DateOfBirth,
		Image:       save.Image,
	}

	return response, nil
}

func (c *customerService) CustomerLogout(ctx *fiber.Ctx) (*fiber.Cookie, error) {
	cookie, err := c.customerService.Logout()
	if err != nil {
		return nil, err
	}

	return cookie, nil
}

func (c *customerService) GetProfile(ctx *fiber.Ctx) (*response.CustomerResponse, error) {
	customer := ctx.Locals("customer").(entity.Customer)

	data, err := c.customerService.Profile(customer)
	if err != nil {
		return nil, err
	}

	cst := &response.CustomerResponse{
		Id:          data.Id,
		Name:        data.Name,
		Username:    data.Username,
		Email:       data.Email,
		Password:    data.Password,
		Phone:       data.Phone,
		Address:     data.Address,
		Gender:      data.Gender,
		DateOfBirth: data.DateOfBirth,
		Image:       data.Image,
	}

	return cst, nil
}

func (c *customerService) CustomerLogin(ctx *fiber.Ctx, input dto.RequestCustomerLogin) (*response.CustomerResponse, string, error) {
	customer, err := c.customerService.Login(input)
	if err != nil {
		return nil, "", err
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(customer.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Subject:   "customer",
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

	respon := &response.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		Username:    customer.Username,
		Email:       customer.Email,
		Password:    customer.Password,
		Phone:       customer.Phone,
		Address:     customer.Address,
		Gender:      customer.Gender,
		DateOfBirth: customer.DateOfBirth,
		Image:       customer.Image,
	}

	return respon, tokens, nil
}

func (c *customerService) CustomerRegister(ctx *fiber.Ctx, input dto.RequestCustomerRegister) (*response.CustomerResponse, error) {
	password, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = password

	image, err := ctx.FormFile("image")
	if err != nil {
		return nil, err
	}

	filename := utils.GenerateImageFile(input.Username, image.Filename)

	if err := ctx.SaveFile(image, filepath.Join(PathImageCustomer, filename)); err != nil {
		return nil, err
	}

	customer := entity.Customer{
		Name:        input.Name,
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		Phone:       input.Phone,
		Address:     input.Address,
		Gender:      input.Gender,
		DateOfBirth: input.DateOfBirth,
		Image:       filename,
	}

	register, err := c.customerService.Register(customer)
	if err != nil {
		return nil, err
	}

	respon := &response.CustomerResponse{
		Id:          register.Id,
		Name:        register.Name,
		Username:    register.Username,
		Email:       register.Email,
		Password:    register.Password,
		Phone:       register.Phone,
		Address:     register.Address,
		Gender:      register.Gender,
		DateOfBirth: register.DateOfBirth,
		Image:       register.Image,
	}

	return respon, nil
}

func NewCustomerService(csr repository.CustomerRepository) CustomerService {
	return &customerService{customerService: csr}
}
