package settings

import (
	controllers "api/the_deck/Controllers"
	repository "api/the_deck/Repository"
	service "api/the_deck/Service"
)

func SetUpServiceAdmin() *controllers.AdminController {
	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository)
	adminController := controllers.NewAdminController(adminService)

	return adminController
}

func SetUpServiceCategory() *controllers.CategoryController {
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	return categoryController
}

func SetUpServiceProduct() *controllers.ProductController {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	producController := controllers.NewProductController(productService)

	return producController
}

func SetUpServiceTable() *controllers.TableController {
	tableRepository := repository.NewTableRepository()
	tableService := service.NewTableService(tableRepository)
	tableController := controllers.NewTableController(tableService)

	return tableController
}

func SetUpServiceCustomer() *controllers.CustomerController {
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository)
	customerController := controllers.NewCustomerController(customerService)

	return customerController
}
