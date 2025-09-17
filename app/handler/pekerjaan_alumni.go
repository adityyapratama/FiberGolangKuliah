package handler

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/service" 
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PekerjaanHandler memegang dependensi ke service-nya.
type PekerjaanHandler struct {
	Svc *service.PekerjaanAlumniService
}

// NewPekerjaanHandler adalah constructor untuk membuat instance handler baru.
func NewPekerjaanHandler(svc *service.PekerjaanAlumniService) *PekerjaanHandler {
	return &PekerjaanHandler{Svc: svc}
}

func (h *PekerjaanHandler) CreatePekerjaanHandler(c *fiber.Ctx) error {
	var req model.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
	}
	newPekerjaan, err := h.Svc.CreatePekerjaanService(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": newPekerjaan})
}

func (h *PekerjaanHandler) GetPekerjaanByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pekerjaan tidak valid"})
	}
	pekerjaan, err := h.Svc.GetPekerjaanByIDService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": pekerjaan})
}

func (h *PekerjaanHandler) GetAllPekerjaanByAlumniIDHandler(c *fiber.Ctx) error {
	alumniID, err := strconv.Atoi(c.Params("alumni_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Alumni ID tidak valid"})
	}
	pekerjaanList, err := h.Svc.GetAllPekerjaanByAlumniIDService(alumniID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": pekerjaanList})
}

func (h *PekerjaanHandler) UpdatePekerjaanHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pekerjaan tidak valid"})
	}
	var req model.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
	}
	updatedPekerjaan, err := h.Svc.UpdatePekerjaanService(id, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": updatedPekerjaan})
}

func (h *PekerjaanHandler) DeletePekerjaanHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pekerjaan tidak valid"})
	}
	if err := h.Svc.DeletePekerjaanService(id); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}



func (h *PekerjaanHandler) GetPekerjaanByIDsajaHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pekerjaan tidak valid"})
	}
	pekerjaan, err := h.Svc.GetPekerjaanByIDService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": pekerjaan})
}

// BARU: GetAllPekerjaanHandler menangani request untuk mendapatkan semua data pekerjaan.
func (h *PekerjaanHandler) GetAllPekerjaansajaHandler(c *fiber.Ctx) error {
	pekerjaanList, err := h.Svc.GetAllPekerjaansajaService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": pekerjaanList})
}