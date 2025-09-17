package route

import (
	"fiber-golang-kuliah/app/handler"
	"fiber-golang-kuliah/middleware"

	"github.com/gofiber/fiber/v2"
)


func SetupPekerjaanRoutes(app fiber.Router, pekerjaanHandler *handler.PekerjaanHandler) {
	pekerjaan := app.Group("/pekerjaan")

	
	pekerjaan.Get("/", pekerjaanHandler.GetAllPekerjaansajaHandler)
	pekerjaan.Get("/:id", pekerjaanHandler.GetPekerjaanByIDHandler)
	pekerjaan.Get("/alumni/:alumni_id", pekerjaanHandler.GetAllPekerjaanByAlumniIDHandler)

	
	pekerjaan.Post("/", middleware.AdminOnly(), pekerjaanHandler.CreatePekerjaanHandler)
	pekerjaan.Put("/:id", middleware.AdminOnly(), pekerjaanHandler.UpdatePekerjaanHandler)
	pekerjaan.Delete("/:id", middleware.AdminOnly(), pekerjaanHandler.DeletePekerjaanHandler)
}