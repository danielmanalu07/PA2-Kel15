package main

import (
	database "api/the_deck/Database"
	migrations "api/the_deck/Database/Migrations"
	routes "api/the_deck/Routes"
	settings "api/the_deck/Settings"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()
	migrations.Migration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	adminController := settings.SetUpServiceAdmin()
	categoryController := settings.SetUpServiceCategory()
	productController := settings.SetUpServiceProduct()
	tableController := settings.SetUpServiceTable()
	customerController := settings.SetUpServiceCustomer()

	routes.RouteAdmin(app, adminController)
	routes.RouteCategory(app, categoryController)
	routes.RouteProduct(app, productController)
	routes.RouteTable(app, tableController)
	routes.RouteCustomer(app, customerController)

	err := app.Listen(":8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
