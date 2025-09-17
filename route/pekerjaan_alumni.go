package route

import (
	"fiber-golang-kuliah/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupPekerjaanAlumniRoutes(app *fiber.App, h *handler.PekerjaanHandler) {
	api := app.Group("/api")

	pekerjaan := api.Group("/pekerjaan")

	pekerjaan.Get("/", h.GetAllPekerjaansajaHandler)
	// Path diubah menjadi "/alumni/:alumni_id" agar sesuai dengan handler-nya.
	pekerjaan.Get("/alumni/:alumni_id", h.GetAllPekerjaanByAlumniIDHandler)

	pekerjaan.Post("/", h.CreatePekerjaanHandler)
	pekerjaan.Get("/:id", h.GetPekerjaanByIDHandler)
	pekerjaan.Get("/:id", h.GetPekerjaanByIDsajaHandler)
	pekerjaan.Put("/:id", h.UpdatePekerjaanHandler)
	pekerjaan.Delete("/:id", h.DeletePekerjaanHandler)
}