package route

import (
	"fiber-golang-kuliah/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAlumniRoutes(app *fiber.App, h *handler.AlumniHandler){
	api := app.Group("/api")

	alumni := api.Group("/alumni")

	alumni.Get("/", h.GetAllAlumniHandler)
	alumni.Post("/", h.CreateAlumniHandler)
	alumni.Get("/:id", h.GetAlumniByIDHandler)
	alumni.Put("/:id", h.UpdateAlumniHandler)
	alumni.Delete("/:id",h.DeleteAlumniHandler)
	

}