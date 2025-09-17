package service

import (
	"errors"
	"fiber-golang-kuliah/app/model"
	"fiber-golang-kuliah/app/repository"
	"fiber-golang-kuliah/utils"
)

type AuthService struct {
	Repo *repository.AuthRepository
}

func NewAuthService( repo *repository.AuthRepository) *AuthService{
	return &AuthService{Repo : repo}
}

func (s *AuthService) LoginService(req model.LoginRequest) (*model.User, string, error){
	if req.Username == "" || req.Password == "" {
		return nil, "", errors.New("username dan password harus diisi")
	}

	user, passwordHash, err := s.Repo.GetUserByUsername(req.Username)

	if err != nil {

		return nil, "", errors.New("username belum ada")
	}


	if !utils.CheckPassword(req.Password, passwordHash) {
		return nil, "", errors.New("username atau password salah")
	}


	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, "", errors.New("gagal membuat token")
	}

	return user, token, nil
}

func (s *AuthService) RegisterService(req model.RegisterRequest) (*model.User, error) {
    
    if req.Username == "" || req.Email == "" || req.Password == "" {
        return nil, errors.New("username, email, dan password wajib diisi")
    }

    
    passwordHash, err := utils.HashPassword(req.Password)
    if err != nil {
        return nil, errors.New("gagal memproses password")
    }

    
    if req.Role == "" {
        req.Role = "user"
    }

    
    user, err := s.Repo.CreateUser(req, passwordHash)
    if err != nil {
    
        return nil, errors.New("gagal mendaftarkan pengguna, username atau email mungkin sudah digunakan")
    }

    return user, nil
}