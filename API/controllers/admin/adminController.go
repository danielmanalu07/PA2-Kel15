package controllers

import (
	"fmt"
	"pa2/database"
	"pa2/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const AdminSecretKey = "admin_secret"

func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	admin := models.Admin{
		Username: data["username"],
		Password: string(password),
	}

	if err := admin.ValidateAdmin(); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			var validationErrors []string
			for _, e := range validateErr {
				validationErrors = append(validationErrors, e.Error())
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": validationErrors,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	database.DB.Create(&admin)
	return ctx.JSON(admin)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var admin models.Admin

	database.DB.Where("username = ?", data["username"]).First(&admin)

	if admin.Username != data["username"] {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := admin.ValidateAdmin(); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			var validationErrors []string
			for _, e := range validateErr {
				validationErrors = append(validationErrors, e.Error())
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": validationErrors,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.Id)),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Subject:   "admin",
	})

	token, err := claims.SignedString([]byte(AdminSecretKey))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "couldn't login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwtAdmin",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
		Secure:   true,
	}

	ctx.Cookie(&cookie)

	if fiber.StatusOK == 200 {
		fmt.Println("Has been login")
	}

	return ctx.JSON(fiber.Map{
		"message": "login success",
	})
}

func Profile(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwtAdmin")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(AdminSecretKey), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var admin models.Admin

	database.DB.Where("id = ?", claims.Issuer).First(&admin)

	return ctx.JSON(admin)
}

func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwtAdmin"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Minute)
	cookie.HTTPOnly = true
	cookie.Secure = true

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
