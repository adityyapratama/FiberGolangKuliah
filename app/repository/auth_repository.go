package repository

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
)

type AuthRepository struct {
	DB *sql.DB

}

func NewAuthRepository (db *sql.DB) *AuthRepository {
return &AuthRepository{DB:db}
}

func (r *AuthRepository) GetUserByUsername(username string) (*model.User, string, error) {
	var user model.User
	var passwordHash string
	query := "SELECT id, username, email, password_hash, role, created_at FROM users WHERE username = $1"
	
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID, 
		&user.Username, 
		&user.Email, 
		&passwordHash, 
		&user.Role, 
		&user.CreatedAt,
	)

	if err != nil {
		return nil, "", err // Mengembalikan error jika user tidak ditemukan atau ada masalah lain.
	}
	
	return &user, passwordHash, nil
}

func (r *AuthRepository) CreateUser( req model.RegisterRequest, passwordHash string) (*model.User, error) {
	var user model.User
	query := `INSERT INTO users (username, email, password_hash, role) 
              VALUES ($1, $2, $3, $4) 
              RETURNING id, username, email, role, created_at`
	
	err := r.DB.QueryRow(query,req.Username, req.Email, passwordHash, req.Role).Scan(
		&user.ID, 
		&user.Username, 
		&user.Email, 
		&user.Role, 
		&user.CreatedAt,
	)

	if err != nil {
        return nil, err
    }
    return &user, nil
}

