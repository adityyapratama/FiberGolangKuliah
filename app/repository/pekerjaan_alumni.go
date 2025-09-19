package repository

import (
	"database/sql"
	"fiber-golang-kuliah/app/model"
	"fmt"
	"log"
	"time"
)

type PekerjaanAlumniRepository struct {
	DB *sql.DB
}

func NewPekerjaanAlumniRepository(db *sql.DB) *PekerjaanAlumniRepository{
	return &PekerjaanAlumniRepository{DB: db}

}


func (r *PekerjaanAlumniRepository) CreatePekerjaan(req model.CreatePekerjaanRequest, tglMulai time.Time, tglSelesai *time.Time) (*model.Pekerjaan, error) {
	var id int
	query := `INSERT INTO pekerjaan_alumni (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	err := r.DB.QueryRow(query, req.AlumniID, req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri, req.LokasiKerja, req.GajiRange, tglMulai, tglSelesai, req.StatusPekerjaan, req.DeskripsiPekerjaan, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.GetPekerjaanByID(id)
}


func (r *PekerjaanAlumniRepository) GetPekerjaanByID(id int) (*model.Pekerjaan, error) {
	var p model.Pekerjaan
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
			  FROM pekerjaan_alumni WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}


func (r *PekerjaanAlumniRepository) GetAllPekerjaanByAlumniID(alumniID int) ([]model.Pekerjaan, error) {
	var pekerjaanList []model.Pekerjaan
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
			  FROM pekerjaan_alumni WHERE alumni_id = $1 ORDER BY tanggal_mulai_kerja DESC`
	rows, err := r.DB.Query(query, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}
	return pekerjaanList, nil
}


func (r *PekerjaanAlumniRepository) UpdatePekerjaan(id int, req model.UpdatePekerjaanRequest, tglMulai time.Time, tglSelesai *time.Time) (*model.Pekerjaan, error) {
	query := `UPDATE pekerjaan_alumni SET 
				nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, gaji_range=$5, 
				tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=$10
			  WHERE id=$11`
	result, err := r.DB.Exec(query, req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri, req.LokasiKerja, req.GajiRange, tglMulai, tglSelesai, req.StatusPekerjaan, req.DeskripsiPekerjaan, time.Now(), id)
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return r.GetPekerjaanByID(id)
}


func (r *PekerjaanAlumniRepository) DeletePekerjaan(id int) error {
	result, err := r.DB.Exec("DELETE FROM pekerjaan_alumni WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}





func (r *PekerjaanAlumniRepository) GetAllPekerjaan() ([]model.Pekerjaan, error) {
	var pekerjaanList []model.Pekerjaan
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
			  FROM pekerjaan_alumni ORDER BY created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}
	return pekerjaanList, nil
}



func (r *PekerjaanAlumniRepository) CountPekerjaan(search string) (int, error) {
    var total int
    query := `SELECT COUNT(id) FROM pekerjaan_alumni WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1`
    // Tambahkan tanda '&' sebelum variabel 'total'
    err := r.DB.QueryRow(query, "%"+search+"%").Scan(&total)
    return total, err
}

func (r *PekerjaanAlumniRepository) GetAllPekerjaanRepo(search, sortBy, order string, limit, offset int) ([]model.Pekerjaan, error) {
	var pekerjaanlist []model.Pekerjaan
	log.Printf("Repository: Menerima parameter search = '%s'", search)
	query := fmt.Sprintf(`
		SELECT id, alumni_id,nama_perusahaan,posisi_jabatan,bidang_industri,lokasi_kerja,gaji_range,tanggal_mulai_kerja,tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan,created_at,updated_at 
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
		`, sortBy, order)
	
	rows, err := r.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Pekerjaan
		
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, 
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, 
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		pekerjaanlist = append(pekerjaanlist, p)
	}
	return pekerjaanlist, nil
}
