package main

import (
	"admin/database"
	"admin/database/migration"
	"admin/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	migration.Migration()
	// seeders.SeederData()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	routes.SetUp(app)

	err := app.Listen("172.27.67.46:8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

}
