package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"fiber-golang-kuliah/app/handler"
	"fiber-golang-kuliah/app/repository"
	"fiber-golang-kuliah/app/service"
	"fiber-golang-kuliah/database"
	"fiber-golang-kuliah/route"
)

func main() {
	// 1. MEMUAT KONFIGURASI .ENV
	// Memuat variabel dari file .env di awal aplikasi
	err := godotenv.Load()
	if err != nil {
		// Program tidak berhenti jika .env tidak ada, hanya memberi peringatan.
		// Ini berguna untuk environment produksi di mana variabel diatur langsung di server.
		log.Println("Peringatan: Tidak dapat memuat file .env")
	}

	// 2. KONEKSI DATABASE
	// Memanggil fungsi ConnectDB yang akan membaca DSN dari environment
	db := database.ConnectDB()
	// defer akan memastikan koneksi ditutup saat fungsi main selesai
	defer db.Close()

	// 3. WIRING / PENYAMBUNGAN DEPENDENSI
	// Inisialisasi dari lapisan terdalam (repository) hingga terluar (handler)
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)

	alumniRepo := repository.NewAlumniRepository(db)
	alumniService :=service.NewAlumniService(alumniRepo)
	alumniHandler := handler.NewAlumniHandler(alumniService)

	pekerjaanRepo := repository.NewPekerjaanAlumniRepository(db)
	pekerjaanService :=service.NewPekerjaanAlumniService(alumniRepo, pekerjaanRepo)
	pekerjaanHandler :=handler.NewPekerjaanHandler(pekerjaanService)

	// 4. INISIALISASI APLIKASI FIBER
	app := fiber.New()

	// 5. SETUP RUTE
	// Memanggil fungsi setup rute dan menyuntikkan handler yang sudah lengkap
	route.SetupMahasiswaRoutes(app, mahasiswaHandler)
	route.SetupAlumniRoutes(app, alumniHandler)
	route.SetupPekerjaanAlumniRoutes(app,pekerjaanHandler)

	// 6. JALANKAN SERVER
	// Ambil port dari .env, jika tidak ada, gunakan "3000" sebagai default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server berjalan di port :%s", port)
	log.Fatal(app.Listen(":" + port))
}