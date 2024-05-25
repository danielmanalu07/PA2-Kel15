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
}

func RouteTable(App *fiber.App, tableController *controllers.TableController) {
	table := App.Group("/table")
	table.Post("/create", tableController.TableCreate)
	table.Get("/", tableController.TableGetAll)
	table.Get("/:id", tableController.TableGetById)
	table.Put("/edit/:id", tableController.TableUpdate)
	table.Delete("/delete/:id", tableController.TableDelete)
}

func RouteRequestTable(App *fiber.App, requestTableController *controllers.RequestTableController) {
	requestTable := App.Group("/request-table")
	requestTable.Post("/create", requestTableController.CreateRequestTable)
	requestTable.Get("/", requestTableController.GetAllRequestTables)
	requestTable.Get("/:id", requestTableController.GetRequestTableById)
	requestTable.Put("/edit/:id", requestTableController.UpdateRequestTable)
	requestTable.Delete("/delete/:id", requestTableController.DeleteRequestTable)
}


func RouteCustomer(App *fiber.App, customerController *controllers.CustomerController) {
	customer := App.Group("/customer")
	customer.Post("/register", customerController.CustomerRegister)
	customer.Post("/login", customerController.CustomerLogin)
	customer.Put("/forgot-password", customerController.CustomerForgotPassword)
	customer.Use(middleware.CheckCustomer())
	customer.Get("/profile", customerController.GetProfile)
	customer.Post("/logout", customerController.CustomerLogout)
	customer.Put("/update-profile", customerController.CustomerUpdateProfile)
	customer.Put("/edit-password", customerController.CustomerEditPassword)
}
