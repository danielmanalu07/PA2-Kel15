package middleware

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
	utils "api/the_deck/Utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CheckLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		tokenString := strings.Replace(cookie, "jwt", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.Secret_Key), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		claims := token.Claims.(*jwt.StandardClaims)
		claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = newToken.SignedString([]byte(utils.Secret_Key))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error refreshing token",
			})
		}

		var admin entity.Admin
		if err := database.DB.Where("username = ?", claims.Issuer).First(&admin).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "customer not found",
			})
		}

		c.Locals("admin", admin)

		return c.Next()
	}
}
