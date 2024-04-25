// package routes

// import (
// 	"pa2/controllers"
// 	"pa2/middleware"

// 	"github.com/gofiber/fiber/v2"
// )

// func SetUpAdmin(app *fiber.App) {
// 	adm := app.Group("/admin")

// 	adm.Post("/create", controllers.AdminCreate)
// 	adm.Post("/login", controllers.AdminLogin)
// 	adm.Get("/logout", controllers.AdminLogout)
// 	adm.Use(middleware.AuthRequired())
// 	adm.Get("/profile", controllers.AdminGetProfile)

// 	adm.Post("/category/create", controllers.CreateCategory)
// 	adm.Get("/category", controllers.GetAllCategory)
// 	adm.Get("/category/:id", controllers.GetAllCategoryById)
// 	adm.Put("/category/:id/edit", controllers.UpdateCategory)
// 	adm.Delete("/category/:id/delete", controllers.DeleteCategory)

//		adm.Post("/product/create", controllers.CreateProduct)
//		adm.Get("/products", controllers.GetAllProduct)
//		adm.Get("/product/:id", controllers.GetProductById)
//		adm.Put("/product/:id/edit", controllers.UpdateProduct)
//		adm.Delete("/product/:id/delete", controllers.DeleteProduct)
//	}
package routes

import (
	"pa2/controllers"
	"pa2/middleware"
	"pa2/repository/implements"

	"github.com/gofiber/fiber/v2"
)

func SetUpAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin")
	admin.Post("/create", controllers.AdminCreate)
	admin.Post("/login", controllers.AdminLogin)
	admin.Get("/logout", controllers.AdminLogout)
	admin.Use(middleware.AuthRequired())
	admin.Get("/profile", controllers.AdminGetProfile)
}

func SetUpCategoryRoutes(app *fiber.App) {
	category := app.Group("/category")
	category.Use(middleware.AuthRequired())
	category.Post("/create", controllers.CreateCategory)
	category.Get("/", controllers.GetAllCategory)
	category.Get("/:id", controllers.GetAllCategoryById)
	category.Put("/:id/edit", controllers.UpdateCategory)
	category.Delete("/:id/delete", controllers.DeleteCategory)
}

func SetUpProductRoutes(app *fiber.App) {
	product := app.Group("/product")
	product.Use(middleware.AuthRequired())
	product.Post("/create", controllers.CreateProduct)
	product.Get("/", controllers.GetAllProduct)
	product.Get("/:id", controllers.GetProductById)
	product.Put("/:id/edit", controllers.UpdateProduct)
	product.Delete("/:id/delete", controllers.DeleteProduct)
	app.Static("/image", implements.PathImageProduct)
}
