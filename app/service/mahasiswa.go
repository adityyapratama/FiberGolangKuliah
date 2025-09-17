package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
)

type MahasiswaService struct {
	Repo *repository.MahasiswaRepository
}

func NewMahasiswaService ( repo *repository.MahasiswaRepository) *MahasiswaService{
	return &MahasiswaService{Repo : repo}
}


func(s *MahasiswaService) GetAllMahasiswaService()([]model.Mahasiswa, error){
	return s.Repo.GetAllMahasiswa()
}

func( s*MahasiswaService) GetMahasiswaByIdService(id int)(*model.Mahasiswa, error){
	return s.Repo.GetMahasiswaByID(id)
}

func (s *MahasiswaService) CreateMahasiswaService(req model.CreateMahasiswaRequest)(*model.Mahasiswa, error){
	if req.NIM == "" || req.Nama == "" || req.Jurusan == ""||req.Email==""{
		return nil, errors.New("semua form wajib di isi")
	}
	return s.Repo.CreateMahasiswa(req)
}

func (s *MahasiswaService) UpdateMahasiswaService(id int, req model.UpdateMahasiswaRequest) (*model.Mahasiswa, error){
	if req.Nama == "" || req.Jurusan == ""||req.Email==""{
	return nil, errors.New("nama, jurusan, dan email harus diisi")
}
return s.Repo.UpdateMahasiswa(id , req)
}

func (s *MahasiswaService) DeleteMahasiswaService(id int) error{
	_, err :=s.Repo.GetMahasiswaByID(id)
	if err != nil{
		return  errors.New("Mahasiswa tidak di temukan")

	}
	return s.Repo.DeleteMahasiswa(id)
}