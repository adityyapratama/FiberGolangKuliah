package handler

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlumniHandler struct {
	Svc *service.AlumniService
}

func NewAlumniHandler (svc *service.AlumniService) *AlumniHandler {
	return &AlumniHandler{Svc :svc}
}

func (h *AlumniHandler) GetAllAlumniHandler( c *fiber.Ctx) error {
	AlumniList, err := h.Svc.GetAllAlumniService()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(fiber.Map{"data" : AlumniList})
}


func (h *AlumniHandler) GetAlumniByIDHandler(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "ID tidak valid/ ditemukan"})
	}

	alumni, err := h.Svc.Repo.GetAlumniByID(id)
	if err != nil{
		if err == sql.ErrNoRows{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" :err.Error()})
	}
	return c.JSON(fiber.Map{"data": alumni})
}


func (h *AlumniHandler) CreateAlumniHandler(c *fiber.Ctx) error {
	var req model.CreateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	newAlumni, err := h.Svc.CreateAlumniService(req)
	if err != nil {
		// PERBAIKAN: Mengirim pesan error yang benar dari 'err.Error()'
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": newAlumni})
}



func( h*AlumniHandler) UpdateAlumniHandler(c *fiber.Ctx) error{
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak Valid"})
	}

	var req model.UpdateAlumniRequest
	if err := c.BodyParser(&req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"request body tidak valid"})
	}

	updatedAlumni, err := h.Svc.UpdateAlumniService(id, req)
	if err != nil {
		if err == sql.ErrNoRows{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error" : "mahasiswa tidak ditemukan"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})


	}
	return  c.JSON(fiber.Map{"data" : updatedAlumni})
}




func(h *AlumniHandler ) DeleteAlumniHandler(c *fiber.Ctx) error{
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ID tidak di temukan"})
	}
	
	err = h.Svc.DeleteAlumniService(id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message" : "Alumni  berhasil di hapus"})
}