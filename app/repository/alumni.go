package repository

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"fmt"
	"log"
	"time"
)

type AlumniRepository struct {
	DB *sql.DB
}

func NewAlumniRepository(db *sql.DB) *AlumniRepository{
	return &AlumniRepository{DB:db}
}


func (r *AlumniRepository) GetAllAlumniRepository()([]model.Alumni, error){
	var alumniList []model.Alumni

	rows, err := r.DB.Query("SELECT id, nim, nama, jurusan, angkatan, email, tahun_lulus,created_at, updated_at FROM alumni ORDER BY created_at DESC")
	
	if err != nil{
		return  nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var a model.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.Email, &a.TahunLulus,&a.CreatedAt,&a.UpdatedAt); err!=nil{
				return  nil, err
		}

		alumniList =append(alumniList, a)
	}
	return alumniList, nil;
}

func (r *AlumniRepository) GetAlumniByID(id int)(*model.Alumni, error){
	var a model.Alumni
	row := r.DB.QueryRow("SELECT id, nim, nama, jurusan, angkatan, email, tahun_lulus,created_at, updated_at FROM alumni  WHERE id = $1", id)
	err:= row.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.Email, &a.TahunLulus,&a.CreatedAt,&a.UpdatedAt)

	if err != nil{
		return  nil, err
	}
	return  &a, nil
}


func (r *AlumniRepository) CreateAlumni(req model.CreateAlumniRequest) (*model.Alumni, error) {
	var id int
	query := `INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err := r.DB.QueryRow(query,
		req.NIM,
		req.Nama,
		req.Jurusan,
		req.Angkatan,
		req.TahunLulus,
		req.Email,
		req.NoTelepon,
		req.Alamat,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	
	return r.GetAlumniByID(id)
}


func (r *AlumniRepository) UpdateAlumni(id int, req model.UpdateAlumniRequest) (*model.Alumni, error) {
	
	query := `UPDATE alumni SET 
				nama = $1, 
				jurusan = $2, 
				angkatan = $3, 
				tahun_lulus = $4, 
				email = $5, 
				no_telepon = $6, 
				alamat = $7, 
				updated_at = $8 
              WHERE id = $9`

	result, err := r.DB.Exec(query,
		req.Nama,
		req.Jurusan,
		req.Angkatan,
		req.TahunLulus,
		req.Email,
		req.NoTelepon,
		req.Alamat,
		time.Now(),
		id,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	
	return r.GetAlumniByID(id)
}


func ( r *AlumniRepository) DeleteAlumni(id int) error{
	result, err := r.DB.Exec("DELETE FROM alumni WHERE id = $1", id)
	if err != nil{
		return err
	}

	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return  nil
} 



func (r *AlumniRepository) CountAlumni(search string) (int, error) {
    var total int
    query := `SELECT COUNT(id) FROM alumni WHERE nama ILIKE $1 OR jurusan ILIKE $1`
    
    err := r.DB.QueryRow(query, "%"+search+"%").Scan(&total)
    return total, err
}

// hitung jumlah data
func (r *AlumniRepository) GetAllAlumni(search, sortBy, order string, limit, offset int) ([]model.Alumni, error) {
	var alumniList []model.Alumni
	log.Printf("Repository: Menerima parameter search = '%s'", search)
	query := fmt.Sprintf(`
		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at 
		FROM alumni
		WHERE nama ILIKE $1 OR jurusan ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
		`, sortBy, order)
	
	rows, err := r.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Alumni
		// Tambahkan &a.NoTelepon dan &a.Alamat
		if err := rows.Scan(
			&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
			&a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}
	return alumniList, nil
}