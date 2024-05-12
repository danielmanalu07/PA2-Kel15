package routes

import (
	controllers "service/table/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	rts := app.Group("table")
	rts.Get("/", controllers.GetAllTable)
	rts.Post("/create", controllers.CreateTable)
	rts.Get("/:id", controllers.GetTableById)
	rts.Put("/:id/edit", controllers.UpdateTable)
	rts.Delete("/:id/delete", controllers.DeleteTable)
}
