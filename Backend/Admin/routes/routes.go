package routes

import (
	controllers "admin/Controllers"
	"admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	endpoint := app.Group("/admin")
	endpoint.Post("/login", controllers.LoginAdmin)
	endpoint.Use(middlewares.Middleware())
	endpoint.Get("/profile", controllers.GetProfile)
	endpoint.Post("/logout", controllers.LogouAdmin)
}
