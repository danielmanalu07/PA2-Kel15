package controllers

import (
	"pa2/database"
	"pa2/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateAnnouncement(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["title"] == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	if data["content"] == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	adminID := ctx.Locals("id").(string)

	var admin models.Admin

	database.DB.Where("id = ?", adminID).Find(&admin)

	announcement := models.Announcement{
		Title:   data["title"],
		Content: data["content"],
		AdminID: admin.Id,
	}

	result := database.DB.Create(&announcement)

	if result.Error != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "Error creating announcement",
		})
	}

	return ctx.JSON(announcement)
}

func IndexAnnouncement(ctx *fiber.Ctx) error {
	var announcement []models.Announcement

	database.DB.Find(&announcement)

	if len(announcement) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Announcement Not Found",
		})
	}

	return ctx.JSON(announcement)
}

func ShowAnnouncement(ctx *fiber.Ctx) error {
	announcementIDStr := ctx.Params("id")

	announcementID, err := strconv.Atoi(announcementIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Announcement ID",
		})
	}

	var announcement models.Announcement

	database.DB.Where("id = ?", announcementID).Find(&announcement)

	return ctx.JSON(announcement)
}

func UpdateAnnouncement(ctx *fiber.Ctx) error {
	announcementIDStr := ctx.Params("id")

	announcementID, err := strconv.Atoi(announcementIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Announcement ID",
		})
	}

	var announcement models.Announcement

	database.DB.Where("id = ?", announcementID).Find(&announcement)

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["title"] == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	if data["content"] == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	updates := map[string]interface{}{
		"Title":   data["title"],
		"Content": data["content"],
	}

	result := database.DB.Model(&announcement).Updates(updates)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating announcement",
		})
	}

	return ctx.JSON(announcement)
}

func DeleteAnnouncement(ctx *fiber.Ctx) error {
	announcementIDStr := ctx.Params("id")

	announcementID, err := strconv.Atoi(announcementIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Announcement ID",
		})
	}

	var announcement models.Announcement

	database.DB.Where("id = ?", announcementID).Find(&announcement)

	if announcementID != int(announcement.ID) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Announcement not found",
		})
	}

	if err := database.DB.Delete(&announcement).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't deleted announcement",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Announcement deleted",
	})
}
