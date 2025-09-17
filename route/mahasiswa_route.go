package route

import (
	"fiber-golang-kuliah/app/handler"
	"fiber-golang-kuliah/middleware"

	"github.com/gofiber/fiber/v2"
)


func SetupMahasiswaRoutes(protected fiber.Router, h *handler.MahasiswaHandler) {
	
	mahasiswa := protected.Group("/mahasiswa")

	
	mahasiswa.Get("/", h.GetAllMahasiswaHandler)
	mahasiswa.Get("/:id", h.GetMahasiswaByIDHandler)

	
	mahasiswa.Post("/", middleware.AdminOnly(), h.CreateMahasiswaHandler)
	mahasiswa.Put("/:id", middleware.AdminOnly(), h.UpdateMahasiswaHandler)
	mahasiswa.Delete("/:id", middleware.AdminOnly(), h.DeleteMahasiswaHandler)
}