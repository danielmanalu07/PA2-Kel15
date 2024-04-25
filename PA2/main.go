package main

import (
	"pa2/database"
	"pa2/database/migration"
	"pa2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()

	migration.Migration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	appAdmin := fiber.New()
	routes.SetUpAdminRoutes(appAdmin)
	go func() {
		_ = appAdmin.Listen(":8080")
	}()

	appCategory := fiber.New()
	routes.SetUpCategoryRoutes(appCategory)
	go func() {
		_ = appCategory.Listen(":8002")
	}()

	appProduct := fiber.New()
	routes.SetUpProductRoutes(appProduct)
	go func() {
		_ = appProduct.Listen(":8003")
	}()

	select {}
	// err := app.Listen(":8080")
	// if err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}
