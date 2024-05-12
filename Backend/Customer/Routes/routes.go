package routes

import (
	controllers "Service/Customer/Controllers"
	middlewares "Service/Customer/Middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetUp(App *fiber.App) {
	customer := App.Group("/customer")
	customer.Post("/register", controllers.RegistrationCustomer)
	customer.Post("/login", controllers.LoginCustomer)
	customer.Static("/image", controllers.PathImageCustomer)
	customer.Put("/forgot_password", controllers.ForgotPassword)
	customer.Use(middlewares.CheckLogin())
	customer.Get("/profile", controllers.GetProfile)
	customer.Put("/updateProfile", controllers.UpdateProfile)
	customer.Put("/edit_password", controllers.EditPassword)
	customer.Post("/logout", controllers.CustomerLogout)

}
