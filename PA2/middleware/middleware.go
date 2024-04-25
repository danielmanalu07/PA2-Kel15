package middleware

import (
	"os"
	"pa2/database"
	"pa2/models/entity"
	"pa2/models/response"
	"pa2/repository/implements"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	AdminRepo = implements.NewAdminRepository()
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("admin")
		if cookie == "" {
			return response.UnauthorizedResponse(c)
		}

		tokenString := strings.TrimSpace(cookie)
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			return response.UnauthorizedResponse(c)
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < time.Hour {
				claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
				newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, err = newToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
				if err != nil {
					return response.InternalServerError(c)
				}

				c.Cookie(&fiber.Cookie{
					Name:     "admin",
					Value:    tokenString,
					Expires:  time.Now().Add(time.Hour * 24),
					HTTPOnly: true,
				})
			}

			var admin entity.Admin
			result := database.DB.Where("username = ?", claims.Issuer).First(&admin)
			if result.Error != nil {
				return response.UnauthorizedResponse(c)
			}

			c.Locals("admin", admin)
			return c.Next()
		}

		return response.UnauthorizedResponse(c)
	}
}
