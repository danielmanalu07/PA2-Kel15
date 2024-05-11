package routes

import (
	controllers "cashier/Controllers"
	middleware "cashier/Middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	cashier := app.Group("/cashier")
	cashier.Post("/register", controllers.CashierRegister)
	cashier.Post("/login", controllers.CashierLogin)
	cashier.Use(middleware.Middleware())
	cashier.Get("/profile", controllers.CashierGetProfile)
	cashier.Post("/logout", controllers.CashierLogout)
}
