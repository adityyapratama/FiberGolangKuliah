package handler

import (
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{Svc: svc}
}

func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
	}
	log.Printf("Mencoba login untuk user: '%s'", req.Username)

	user, token, err := h.Svc.LoginService(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}


func (h *AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err :=c.BodyParser(&req); err != nil{
		return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "request body tidak valid",
		})
	}

	user, err :=h.Svc.RegisterService(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
        "message": "Registrasi berhasil",
        "data":    user,
	})


}