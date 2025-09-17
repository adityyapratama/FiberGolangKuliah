package handler

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MahasiswaHandler struct {
	Svc *service.MahasiswaService
}

func NewMahasiswaHandler(svc *service.MahasiswaService) *MahasiswaHandler{
	return &MahasiswaHandler{Svc :svc}
}

func (h *MahasiswaHandler) GetAllMahasiswaHandler( c *fiber.Ctx) error {
	mahasiswaList, err := h.Svc.GetAllMahasiswaService()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(fiber.Map{"data" : mahasiswaList})
}


func (h *MahasiswaHandler) GetMahasiswaByIDHandler(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "ID tidak valid/ ditemukan"})
	}

	mahasiswa, err := h.Svc.GetMahasiswaByIdService(id)
	if err != nil{
		if err == sql.ErrNoRows{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mahasiswa tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" :err.Error()})
	}
	return c.JSON(fiber.Map{"data": mahasiswa})
}

func (h *MahasiswaHandler) CreateMahasiswaHandler(c *fiber.Ctx) error{
	var req model.CreateMahasiswaRequest
	if err := c.BodyParser(&req); err!= nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	NewMahasiswa, err := h.Svc.CreateMahasiswaService(req)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : NewMahasiswa})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": NewMahasiswa})
}


func( h*MahasiswaHandler) UpdateMahasiswaHandler(c *fiber.Ctx) error{
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak Valid"})
	}

	var req model.UpdateMahasiswaRequest
	if err := c.BodyParser(&req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"request body tidak valid"})
	}

	updatedMahasiswa, err := h.Svc.UpdateMahasiswaService(id, req)
	if err != nil {
		if err == sql.ErrNoRows{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error" : "mahasiswa tidak ditemukan"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})


	}
	return  c.JSON(fiber.Map{"data" : updatedMahasiswa})
}

func(h *MahasiswaHandler ) DeleteMahasiswaHandler(c *fiber.Ctx) error{
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ID tidak di temukan"})
	}
	
	err = h.Svc.DeleteMahasiswaService(id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message" : "mahasiswa  berhasil di hapus"})
}