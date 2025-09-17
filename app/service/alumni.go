package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
)

type AlumniService struct {
	Repo *repository.AlumniRepository
}

func NewAlumniService (repo *repository.AlumniRepository) *AlumniService {
	return &AlumniService{Repo : repo}
}


func(a *AlumniService) GetAllAlumniService()([]model.Alumni, error){
	return a.Repo.GetAllAlumniRepository()
}

func(a *AlumniService) GetAlumniService(id int)(*model.Alumni, error){
	return a.Repo.GetAlumniByID(id)
}

func (s *AlumniService) CreateAlumniService(req model.CreateAlumniRequest) (*model.Alumni, error) {
	if req.NIM == "" || req.Nama == "" || req.Jurusan == "" || req.Angkatan == 0 || req.Email == "" || req.TahunLulus == 0 {
		return nil, errors.New("NIM, Nama, Jurusan, Angkatan, Tahun Lulus, dan Email wajib diisi")
	}
	return s.Repo.CreateAlumni(req)

}

func (a *AlumniService) UpdateAlumniService(id int, req model.UpdateAlumniRequest) (*model.Alumni, error){
	if req.NIM == "" || req.Nama == "" || req.Jurusan == "" || req.Angkatan == 0 || req.Email == "" || req.TahunLulus == 0 {
	return nil, errors.New("nama, jurusan, dan email harus diisi")
}
return a.Repo.UpdateAlumni(id , req)
}


func (a *AlumniService) DeleteAlumniService(id int) error{
	_, err :=a.Repo.GetAlumniByID(id)
	if err != nil{
		return  errors.New("Mahasiswa tidak di temukan")

	}
	return a.Repo.DeleteAlumni(id)
}