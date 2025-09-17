package database

import(
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"	
)
func ConnectDB() *sql.DB {
	// Ambil DSN dari environment variable yang bernama DB_DSN
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Gagal ping database:", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL")
	return db
}
// var DB *sql.DB

// func ConnectDB(){
// 	var err error
// 	// connection 
// 	dsn := "host=localhost user=postgres password=password dbname=Mahasiswa port=5433 sslmode=disable"

// 	DB, err = sql.Open("postgres", dsn)
// 	if err != nil{
// 		log.Fatal("GAGAL TERKONEKSI KE DATABASE :", err)
// 	}

// 	// test connection
// 	if err =DB.Ping(); err != nil{
// 		log.Fatal("GAGAK PING KE DATABASE :", err)	
// 	}

// 	fmt.Println("SUKSES TERKONEKSI KE DATABASE")
// 	return DB


// }