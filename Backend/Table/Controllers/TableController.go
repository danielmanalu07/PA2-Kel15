package controllers

import (
	database "service/table/Database"
	"service/table/Models/entity"
	"service/table/Models/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllTable(c *fiber.Ctx) error {
	var table []entity.Table

	result := database.DB.Debug().Find(&table)

	if len(table) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Not Found",
		})
	}

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Couldn't find data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": table,
	})
}

func CreateTable(c *fiber.Ctx) error {
	input := new(response.RequestTableCrate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}
	tables := entity.Table{
		Number:   input.Number,
		Capacity: input.Capacity,
	}

	err := database.DB.Create(&tables).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Couldn't to Create data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": tables,
	})
}

func GetTableById(c *fiber.Ctx) error {
	id := c.Params("id")

	var tables entity.Table

	err := database.DB.Where("id = ?", id).First(&tables).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": tables,
	})
}

func UpdateTable(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(response.RequestTableUpdate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	var tables entity.Table

	err := database.DB.Where("id = ?", id).First(&tables).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Not Found",
		})
	}

	if input.Number != 0 {
		tables.Number = input.Number
	}

	if input.Capacity != 0 {
		tables.Capacity = input.Capacity
	}

	err = database.DB.Save(&tables).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Couldn't to Update data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": tables,
	})
}

func DeleteTable(c *fiber.Ctx) error {
	id := c.Params("id")

	var tables entity.Table

	err := database.DB.Where("id = ?", id).First(&tables).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Not Found",
		})
	}

	if err := database.DB.Debug().Delete(&tables).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Couldn't to Delete Data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Successfully",
	})
}
