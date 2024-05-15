package main

import (
	database "Service/Customer/Database"
	routes "Service/Customer/Routes"
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

	routes.SetUp(app)

	err := app.Listen("192.168.30.215:8006")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
