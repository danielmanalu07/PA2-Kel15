package controllers

import (
	"pa2/database"
	"pa2/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(ctx *fiber.Ctx) error {

	productID := ctx.FormValue("product_id")
	if productID == "" {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "Product is required",
		})
	}

	count := ctx.FormValue("count")
	if count == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Count is required",
		})
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var product models.Product
	database.DB.Find(&product, productIDInt)

	CustomerId := ctx.Locals("id").(string)

	var customer models.Customer
	database.DB.Where("id = ?", CustomerId).Find(&customer)

	booking := models.BookingQueue{
		ProductID:  uint(productIDInt),
		CustomerID: customer.ID,
		Count:      count,
		Taking:     0,
		Status:     0,
	}

	if productIDInt != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	result := database.DB.Create(&booking)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "couldn't create booking",
		})
	}
	return ctx.JSON(booking)
}

func IndexBooking(ctx *fiber.Ctx) error {
	var booking []models.BookingQueue

	customerID := ctx.Locals("id").(string)

	database.DB.Where("customer_id = ?", customerID).Find(&booking)

	if len(booking) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No booking",
		})
	}

	return ctx.JSON(booking)
}

func ShowBooking(ctx *fiber.Ctx) error {
	bookingIDStr := ctx.Params("id")

	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid booking ID",
		})
	}
	customerID := ctx.Locals("id").(string)

	var booking models.BookingQueue

	result := database.DB.Where("id = ? AND customer_id = ?", bookingID, customerID).Find(&booking)

	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return ctx.JSON(booking)
}

func UpdateBooking(ctx *fiber.Ctx) error {
	productID := ctx.FormValue("product_id")
	if productID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Product is required",
		})
	}

	count := ctx.FormValue("count")
	if count == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Count is required",
		})
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	bookingIDStr := ctx.Params("id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid booking ID",
		})
	}

	customerID := ctx.Locals("id").(string)

	var booking models.BookingQueue
	result := database.DB.Where("id = ? AND customer_id = ?", bookingID, customerID).Find(&booking)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	updates := map[string]interface{}{
		"ProductID": productIDInt,
		"Count":     count,
	}

	result = database.DB.Model(&booking).Updates(updates)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating booking",
		})
	}

	database.DB.Find(&booking, bookingID)

	return ctx.JSON(booking)
}

func DeleteBooking(ctx *fiber.Ctx) error {
	bookingIDStr := ctx.Params("id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid booking ID",
		})
	}

	customerID := ctx.Locals("id").(string)

	var booking models.BookingQueue
	result := database.DB.Where("id = ? AND customer_id = ?", bookingID, customerID).Find(&booking)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	if err := database.DB.Delete(&booking).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete booking",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Booking deleted successfully",
	})
}
