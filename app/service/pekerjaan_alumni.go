package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
	
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)


type PekerjaanAlumniService struct {
	AlumniRepo    *repository.AlumniRepository
	PekerjaanRepo *repository.PekerjaanAlumniRepository
}


func NewPekerjaanAlumniService(
	alumniRepo *repository.AlumniRepository,
	pekerjaanRepo *repository.PekerjaanAlumniRepository,
) *PekerjaanAlumniService {
	return &PekerjaanAlumniService{
		AlumniRepo:    alumniRepo,
		PekerjaanRepo: pekerjaanRepo,
	}
}


func parseTanggal(tanggalStr string) (time.Time, error) {
	return time.Parse("2006-01-02", tanggalStr)
}


func (s *PekerjaanAlumniService) CreatePekerjaanService(req model.CreatePekerjaanRequest) (*model.Pekerjaan, error) {
	if req.AlumniID == 0 || req.NamaPerusahaan == "" || req.PosisiJabatan == "" || req.TanggalMulaiKerja == "" {
		return nil, errors.New("Alumni ID, Nama Perusahaan, Posisi, dan Tanggal Mulai Kerja wajib diisi")
	}

	// Memanggil AlumniRepo untuk validasi.
	_, err := s.AlumniRepo.GetAlumniByID(req.AlumniID)
	if err != nil {
		return nil, errors.New("alumni dengan ID tersebut tidak ditemukan")
	}

	tglMulai, err := parseTanggal(req.TanggalMulaiKerja)
	if err != nil {
		return nil, errors.New("format Tanggal Mulai Kerja tidak valid, gunakan YYYY-MM-DD")
	}

	var tglSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		parsed, err := parseTanggal(*req.TanggalSelesaiKerja)
		if err != nil {
			return nil, errors.New("format Tanggal Selesai Kerja tidak valid, gunakan YYYY-MM-DD")
		}
		tglSelesai = &parsed
	}

	// Memanggil PekerjaanRepo untuk membuat data.
	return s.PekerjaanRepo.CreatePekerjaan(req, tglMulai, tglSelesai)
}

func (s *PekerjaanAlumniService) GetPekerjaanByIDService(id int) (*model.Pekerjaan, error) {
	return s.PekerjaanRepo.GetPekerjaanByID(id)
}

func (s *PekerjaanAlumniService) GetAllPekerjaanByAlumniIDService(alumniID int) ([]model.Pekerjaan, error) {
	return s.PekerjaanRepo.GetAllPekerjaanByAlumniID(alumniID)
}

func (s *PekerjaanAlumniService) UpdatePekerjaanService(id int, req model.UpdatePekerjaanRequest) (*model.Pekerjaan, error) {
	if req.NamaPerusahaan == "" || req.PosisiJabatan == "" || req.TanggalMulaiKerja == "" {
		return nil, errors.New("Nama Perusahaan, Posisi, dan Tanggal Mulai Kerja wajib diisi")
	}

	tglMulai, err := parseTanggal(req.TanggalMulaiKerja)
	if err != nil {
		return nil, errors.New("format Tanggal Mulai Kerja tidak valid, gunakan YYYY-MM-DD")
	}

	var tglSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		parsed, err := parseTanggal(*req.TanggalSelesaiKerja)
		if err != nil {
			return nil, errors.New("format Tanggal Selesai Kerja tidak valid, gunakan YYYY-MM-DD")
		}
		tglSelesai = &parsed
	}
	
	return s.PekerjaanRepo.UpdatePekerjaan(id, req, tglMulai, tglSelesai)
}

func (s *PekerjaanAlumniService) DeletePekerjaanService(id int) error {
	return s.PekerjaanRepo.DeletePekerjaan(id)
}



func (s *PekerjaanAlumniService) GetPekerjaansajaByIDService(id int) (*model.Pekerjaan, error) {
	return s.PekerjaanRepo.GetPekerjaanByID(id)
}

func (s *PekerjaanAlumniService) GetAllPekerjaansajaService() ([]model.Pekerjaan, error) {
	return s.PekerjaanRepo.GetAllPekerjaan()
}


func (s *PekerjaanAlumniService) GetAllPekerjaanAlumniServiceSorting(c *fiber.Ctx) (*model.PekerjaanResponse, error) {
    // mengambil query parameter dari URL
    page, _ := strconv.Atoi(c.Query("page", "1"))
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    sortBy := c.Query("sortBy", "id")
    order := c.Query("order", "asc")
    search := c.Query("search", "")


    
    sortByWhitelist := map[string]bool{"id": true, "nama_perusahaan": true, "posisi_jabatan": true, "lokasi_kerja": true}
    if !sortByWhitelist[sortBy] {
        sortBy = "id" 
    }
    if strings.ToLower(order) != "desc" {
        order = "asc"
    }

    
    offset := (page - 1) * limit

    
    alumni, err := s.PekerjaanRepo.GetAllPekerjaanRepo(search, sortBy, order, limit, offset)
    if err != nil {
        return nil, err
    }
    total, err := s.PekerjaanRepo.CountPekerjaan(search)
    if err != nil {
        return nil, err
    }

    
    response := &model.PekerjaanResponse{
        Data: alumni,
        Meta: model.MetaInfo{
            Page:   page,
            Limit:  limit,
            Total:  total,
            Pages:  (total + limit - 1) / limit, // Kalkulasi total halaman
            SortBy: sortBy,
            Order:  order,
            Search: search,
        },
    }
    return response, nil
}