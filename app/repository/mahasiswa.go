package repository

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"time"
)

type MahasiswaRepository struct {
	DB *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) *MahasiswaRepository{
	return &MahasiswaRepository{DB:db}
}

// get all
func (r *MahasiswaRepository)GetAllMahasiswa()([]model.Mahasiswa, error){
	var mahasiswaList []model.Mahasiswa

	rows, err := r.DB.Query("SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa ORDER BY created_at DESC")
	
	if err != nil{
		return  nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var m model.Mahasiswa
		if err := rows.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt); err!=nil{
				return  nil, err
		}

		mahasiswaList =append(mahasiswaList, m)
	}
	return mahasiswaList, nil;
}

// GetByID
func (r *MahasiswaRepository) GetMahasiswaByID(id int)(*model.Mahasiswa, error){
	var m model.Mahasiswa
	row := r.DB.QueryRow("SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa WHERE id = $1", id)
	err:= row.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt)

	if err != nil{
		return  nil, err
	}
	return  &m, nil
}



// createMhs
func (r *MahasiswaRepository) CreateMahasiswa(req model.CreateMahasiswaRequest) (*model.Mahasiswa, error){
	var id int
	err := r.DB.QueryRow(
		"INSERT INTO mahasiswa (nim, nama, jurusan, angkatan, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		req.NIM, req.Nama, req.Jurusan, req.Angkatan, req.Email, time.Now(), time.Now(),
	).Scan(&id)
	if err != nil{
		return	nil, err
	}

	return r.GetMahasiswaByID(id)

}

// updateMhs
func (r *MahasiswaRepository) UpdateMahasiswa(id int, req model.UpdateMahasiswaRequest) (*model.Mahasiswa, error){
	result, err := r.DB.Exec(
		"UPDATE mahasiswa SET nama = $1, jurusan = $2, angkatan = $3, email = $4, updated_at = $5 WHERE id = $6",
		req.Nama, req.Jurusan, req.Angkatan, req.Email, time.Now(), id,
	)
	if err != nil{
		return  nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0{
		return nil, sql.ErrNoRows
	}
	return  r.GetMahasiswaByID(id)

}

// delete
func ( r *MahasiswaRepository) DeleteMahasiswa (id int) error{
	result, err := r.DB.Exec("DELETE FROM mahasiswa WHERE id = $1", id)
	if err != nil{
		return err
	}

	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return  nil
} 

