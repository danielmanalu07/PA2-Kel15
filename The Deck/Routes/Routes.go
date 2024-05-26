package routes

import (
	controllers "api/the_deck/Controllers"
	middleware "api/the_deck/Middleware"
	service "api/the_deck/Service"

	"github.com/gofiber/fiber/v2"
)

func RouteAdmin(App *fiber.App, adminController *controllers.AdminController) {
	admin := App.Group("/admin")
	admin.Post("/login", adminController.AdminLogin)
	admin.Use(middleware.CheckLogin())
	admin.Get("/profile", adminController.GetProfile)
	admin.Post("/logout", adminController.LogoutAdmin)
}

func RouteCategory(App *fiber.App, categoryController *controllers.CategoryController) {
	category := App.Group("/category")
	category.Post("/create", categoryController.CategoryCreate)
	category.Get("/", categoryController.CategoryGetAll)
	category.Get("/:id", categoryController.CategoryGetById)
	category.Put("/edit/:id", categoryController.CategoryUpdate)
	category.Delete("/delete/:id", categoryController.CategoryDelete)
}

func RouteProduct(App *fiber.App, productController *controllers.ProductController) {
	product := App.Group("/product")
	product.Static("/image", service.PathImageProduct)
	product.Post("/create", productController.ProductCreate)
	product.Get("/", productController.ProductGetAll)
	product.Get("/:id", productController.ProductGetById)
	product.Put("/edit/:id", productController.ProductUpdate)
	product.Delete("/delete/:id", productController.ProductDelete)
	product.Get("/category/:cat", productController.ProductGetByCategory)
}

func RouteTable(App *fiber.App, tableController *controllers.TableController) {
	table := App.Group("/table")
	table.Post("/create", tableController.TableCreate)
	table.Get("/", tableController.TableGetAll)
	table.Get("/:id", tableController.TableGetById)
	table.Put("/edit/:id", tableController.TableUpdate)
	table.Delete("/delete/:id", tableController.TableDelete)
}

func RouteCustomer(App *fiber.App, customerController *controllers.CustomerController) {
	customer := App.Group("/customer")
	customer.Post("/register", customerController.CustomerRegister)
	customer.Post("/login", customerController.CustomerLogin)
	customer.Static("/image", service.PathImageCustomer)
	customer.Put("/forgot-password", customerController.CustomerForgotPassword)
	customer.Use(middleware.CheckCustomer())
	customer.Get("/profile", customerController.GetProfile)
	customer.Post("/logout", customerController.CustomerLogout)
	customer.Put("/update-profile", customerController.CustomerUpdateProfile)
	customer.Put("/edit-password", customerController.CustomerEditPassword)
}

func RouteCart(App *fiber.App, cartController *controllers.CartController) {
	cart := App.Group("/cart")
	cart.Use(middleware.CheckCustomer())
	cart.Post("/add", cartController.AddItemCart)
	cart.Get("/myCart", cartController.GetItemMyCart)
	cart.Delete("/delete/:id", cartController.DeleteMyCart)
	cart.Put("/edit/:id", cartController.UpdateQuantity)
}

func RouteOrder(App *fiber.App, orderController *controllers.OrderController) {
	order := App.Group("/order")
	order.Get("/", orderController.GetAllOrder)
	order.Use(middleware.CheckCustomer())
	order.Post("/create", orderController.CustomerCreateOrder)
}
