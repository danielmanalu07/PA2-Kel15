package routes

import (
	controllers "pa2/controllers/admin"
	"pa2/middleware"

	cashierController "pa2/controllers/cashier"

	fiber "github.com/gofiber/fiber/v2"

	customerController "pa2/controllers/customer"
)

func Setup(app *fiber.App) {
	admin := app.Group("/admin")

	//Auth Admin
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Use(middleware.RequiredLoginAdmin)
	admin.Get("/profile", controllers.Profile)
	admin.Post("/logout", controllers.Logout)

	//Manage Cashier
	admin.Post("/cashier", controllers.CreateCashier)
	admin.Get("/cashier/index", controllers.IndexCashier)
	admin.Get("/cashier/show/:id", controllers.ShowCashier)
	admin.Put("/cashier/update/:id", controllers.UpdateCashier)
	admin.Delete("/cashier/delete/:id", controllers.DeleteCashier)

	//Manage Category with Admin
	admin.Post("/category", controllers.CreateCategory)
	admin.Get("/category/index", controllers.IndexCategory)
	admin.Get("/category/show/:id", controllers.ShowCategory)
	admin.Put("/category/update/:id", controllers.UpdateCategory)
	admin.Delete("/category/delete/:id", controllers.DeleteCategory)

	//Manage Tabel with Admin
	admin.Post("/table", controllers.CreateTable)
	admin.Get("/table/index", controllers.IndexTable)
	admin.Get("/table/show/:id", controllers.ShowTable)
	admin.Put("/table/update/:id", controllers.UpdateTable)
	admin.Delete("/table/delete/:id", controllers.DeleteTable)

	//Manage Product with Admin
	admin.Post("/product", controllers.CreateProduct)
	admin.Get("/product/index", controllers.IndexProduct)
	admin.Get("/product/show/:id", controllers.ShowProduct)
	admin.Put("/product/update/:id", controllers.UpdateProduct)
	admin.Delete("/product/delete/:id", controllers.DeleteProduct)

	//manage Announcement with Admin
	admin.Post("/announcement", controllers.CreateAnnouncement)
	admin.Get("/announcement/index", controllers.IndexAnnouncement)
	admin.Get("/announcement/show/:id", controllers.ShowAnnouncement)
	admin.Put("/announcement/update/:id", controllers.UpdateAnnouncement)
	admin.Delete("/announcement/delete/:id", controllers.DeleteAnnouncement)

	//Auth Cashier
	cashier := app.Group("/cashier")
	cashier.Post("/login", cashierController.Login)
	cashier.Use(middleware.RequiredLoginCashier)
	cashier.Get("/profile", cashierController.Profile)
	cashier.Post("/logout", cashierController.Logout)
	//Approval Cashier
	cashier.Put("/approvebooking/:id", cashierController.ApproveBooking)
	cashier.Get("/getbookings", cashierController.GetBooking)
	cashier.Put("/rejectbooking/:id", cashierController.RejectBooking)

	//Auth Customer
	customer := app.Group("/customer")
	customer.Post("/register", customerController.RegisterCustomer)
	customer.Post("/login", customerController.LoginCustomer)
	customer.Use(middleware.RequiredLoginCustomer)
	customer.Get("/profile", customerController.Profile)
	customer.Post("/logout", customerController.Logout)

	//BookingQueue Customer
	customer.Post("/booking", customerController.CreateBooking)
	customer.Get("/booking/index", customerController.IndexBooking)
	customer.Get("/booking/show/:id", customerController.ShowBooking)
	customer.Put("/booking/update/:id", customerController.UpdateBooking)
	customer.Delete("/booking/delete/:id", customerController.DeleteBooking)

	//Proof Of Payment Customer
	customer.Post("/proof-of-payment", customerController.SendPayment)

}
