package routes

import (
	controllers "pa2/controllers/admin"
	"pa2/middleware"

	cashierController "pa2/controllers/cashier"

	fiber "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	admin := app.Group("/admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Use(middleware.RequiredLogin)
	admin.Get("/profile", controllers.Profile)
	admin.Post("/logout", controllers.Logout)

	admin.Post("/cashier", controllers.CreateCashier)
	admin.Get("/cashier/index", controllers.IndexCashier)
	admin.Get("/cashier/show/:id", controllers.ShowCashier)
	admin.Put("/cashier/update/:id", controllers.UpdateCashier)
	admin.Delete("/cashier/delete/:id", controllers.DeleteCashier)

	cashier := app.Group("/cashier")
	cashier.Post("/login", cashierController.Login)
	cashier.Use(middleware.RequiredLogin)
	cashier.Get("/profile", cashierController.Profile)
	cashier.Post("/logout", cashierController.Logout)
}
