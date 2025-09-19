package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
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

func (a *AlumniService) GetAllAlumniServiceSorting(c *fiber.Ctx) (*model.AlumniResponse, error) {
    // mengambil query parameter dari URL
    page, _ := strconv.Atoi(c.Query("page", "1"))
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    sortBy := c.Query("sortBy", "id")
    order := c.Query("order", "asc")
    search := c.Query("search", "")

	 log.Printf("Service: Menerima parameter search = '%s'", search)

    
    sortByWhitelist := map[string]bool{"id": true, "nama": true, "angkatan": true, "tahun_lulus": true}
    if !sortByWhitelist[sortBy] {
        sortBy = "id" 
    }
    if strings.ToLower(order) != "desc" {
        order = "asc"
    }

    
    offset := (page - 1) * limit

    
    alumni, err := a.Repo.GetAllAlumni(search, sortBy, order, limit, offset)
    if err != nil {
        return nil, err
    }
    total, err := a.Repo.CountAlumni(search)
    if err != nil {
        return nil, err
    }

    
    response := &model.AlumniResponse{
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