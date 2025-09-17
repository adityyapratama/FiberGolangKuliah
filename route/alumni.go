package route

import (
	"fiber-golang-kuliah/app/handler"
	"fiber-golang-kuliah/middleware"

	"github.com/gofiber/fiber/v2"
)


func SetupAlumniRoutes(protected fiber.Router, alumniHandler *handler.AlumniHandler) {
	alumni := protected.Group("/alumni")

	
	alumni.Get("/", alumniHandler.GetAllAlumniHandler)
	alumni.Get("/:id", alumniHandler.GetAlumniByIDHandler)

	
	alumni.Post("/", middleware.AdminOnly(), alumniHandler.CreateAlumniHandler)
	alumni.Put("/:id", middleware.AdminOnly(), alumniHandler.UpdateAlumniHandler)
	alumni.Delete("/:id", middleware.AdminOnly(), alumniHandler.DeleteAlumniHandler)
}