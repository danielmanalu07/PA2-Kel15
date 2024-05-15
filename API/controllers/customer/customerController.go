package controllers

import (
	"fmt"
	"os"
	"pa2/database"
	"pa2/models"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const FileCustomer = "./FileCustomer"

const CustomerSecretKey = "customer_secret"

var BlacklistedTokens = make(map[string]bool)

func init() {
	if _, err := os.Stat(FileCustomer); os.IsNotExist(err) {
		os.Mkdir(FileCustomer, os.ModePerm)
	}
}

func RegisterCustomer(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	if username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username is required.",
		})
	}

	phone := ctx.FormValue("phone")
	if phone == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone is required.",
		})
	}

	address := ctx.FormValue("address")
	if address == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Address is required.",
		})
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Photo is required",
		})
	}

	filename := fmt.Sprintf("Customer_%s%s", username, filepath.Ext(photo.Filename))

	if err := ctx.SaveFile(photo, filepath.Join(FileCustomer, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save photo",
		})
	}

	password := ctx.FormValue("password")
	if password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password is required.",
		})
	}

	confirmPassword := ctx.FormValue("confirm_password")
	if confirmPassword == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Confirm password is required.",
		})
	}

	if password != confirmPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Passwords and Confirm_Password do not match.",
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate password hash.",
		})
	}

	customer := models.Customer{
		Username: username,
		Phone:    phone,
		Address:  address,
		Photo:    filename,
		Password: string(hashPassword),
	}

	result := database.DB.Create(&customer)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create customer.",
			"error":   result.Error.Error(),
		})
	}

	customer.Password = ""

	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func LoginCustomer(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var customer models.Customer

	database.DB.Where("username = ?", data["username"]).Find(&customer)

	if customer.Username != data["username"] {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Customer not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(data["password"])); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(customer.ID)),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Subject:   "customer",
	})

	token, err := claims.SignedString([]byte(CustomerSecretKey))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't Login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwtCustomer",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "Login Successfully",
	})
}

func Profile(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwtCustomer")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(CustomerSecretKey), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if BlacklistedTokens[cookie] {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token has been revoked",
		})
	}

	RemoveExpiredToken()

	claims := token.Claims.(*jwt.StandardClaims)

	var customer models.Customer

	database.DB.Where("id = ?", claims.Issuer).First(&customer)

	return ctx.JSON(customer)
}

func RemoveExpiredToken() {
	for token, expired := range BlacklistedTokens {
		if expired {
			delete(BlacklistedTokens, token)
		}
	}
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwtCustomer",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute),
		HTTPOnly: true,
	}

	BlacklistedTokens[cookie.Value] = true
	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "Logout successfully",
	})
}
