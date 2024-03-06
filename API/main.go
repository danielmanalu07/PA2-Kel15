package main

import (
	"log"
	"pa2/database"
	"pa2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
	}))

	routes.Setup(app)

	err := app.Listen(":8001")
	if err != nil {
		log.Fatal(err)
	}
}
