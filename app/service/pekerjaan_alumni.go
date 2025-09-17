package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
	"time"
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