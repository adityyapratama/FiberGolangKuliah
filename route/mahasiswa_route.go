package route

import (
	"fiber-golang-kuliah/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupMahasiswaRoutes(app *fiber.App, h *handler.MahasiswaHandler){
	api := app.Group("/api")

	mahasiswa := api.Group("/mahasiswa")

	mahasiswa.Get("/", h.GetAllMahasiswaHandler)
	mahasiswa.Get("/:id", h.GetMahasiswaByIDHandler)
	mahasiswa.Post("/", h.CreateMahasiswaHandler)
	mahasiswa.Put("/:id", h.UpdateMahasiswaHandler)
	mahasiswa.Delete("/:id", h.DeleteMahasiswaHandler)

}