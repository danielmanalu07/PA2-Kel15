package main

import (
	database "Service/Category/Database"
	routes "Service/Category/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	routes.SetUpCategory(app)
	err := app.Listen(":8002")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
