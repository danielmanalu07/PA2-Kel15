package controllers

import (
	"pa2/database"
	"pa2/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const CashierSecretKey = "cashier_secret"

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var cashier models.Cashier

	database.DB.Where("username = ? ", data["username"]).First(&cashier)

	if cashier.Username != data["username"] {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Account not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cashier.Password), []byte(data["password"])); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(cashier.Id)),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Subject:   "cashier",
	})

	token, err := claims.SignedString([]byte(CashierSecretKey))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't Login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwtCashier",
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
	cookie := ctx.Cookies("jwtCashier")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(CashierSecretKey), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var cashier models.Cashier

	database.DB.Where("id = ?", claims.Issuer).First(&cashier)

	return ctx.JSON(cashier)
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwtCashier",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "Logout successfully",
	})
}

func GetBooking(ctx *fiber.Ctx) error {
	var booking []models.BookingQueue

	database.DB.Find(&booking)

	if len(booking) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Booking not found",
		})
	}

	return ctx.JSON(booking)
}

func ApproveBooking(ctx *fiber.Ctx) error {
	bookingIDStr := ctx.Params("id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid booking ID",
		})
	}

	var booking models.BookingQueue
	result := database.DB.Where("id = ?", bookingID).Find(&booking)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	if booking.Status == 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Booking already approved",
		})
	}

	if booking.Status == 2 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Booking already rejected",
		})
	}

	update := map[string]interface{}{
		"Status": 1,
	}

	result = database.DB.Model(&booking).Updates(update)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error approving booking",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Booking approved successfully",
	})
}

func RejectBooking(ctx *fiber.Ctx) error {
	bookingIDStr := ctx.Params("id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid booking ID",
		})
	}

	var booking models.BookingQueue
	result := database.DB.Where("id = ?", bookingID).Find(&booking)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	if booking.Status == 2 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Booking already rejected",
		})
	}

	if booking.Status == 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Booking already approved",
		})
	}

	update := map[string]interface{}{
		"Status": 2,
	}

	result = database.DB.Model(&booking).Updates(update)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error rejecting booking",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Booking rejected successfully",
	})
}
