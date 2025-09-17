package route

import (
	"fiber-golang-kuliah/app/handler"

	"github.com/gofiber/fiber/v2"
)


func SetupAuthRoutes(api fiber.Router, authHandler *handler.AuthHandler) {
	api.Post("/login", authHandler.LoginHandler)
	api.Post("/register", authHandler.RegisterHandler)
}