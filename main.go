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
	"fiber-golang-kuliah/middleware"
	"fiber-golang-kuliah/route"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak dapat memuat file .env")
	}


	db := database.ConnectDB()
	defer db.Close()




	authRepo := repository.NewAuthRepository(db)
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	pekerjaanRepo := repository.NewPekerjaanAlumniRepository(db)


	authService := service.NewAuthService(authRepo)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	alumniService := service.NewAlumniService(alumniRepo)
	pekerjaanService := service.NewPekerjaanAlumniService(alumniRepo, pekerjaanRepo) // Service ini butuh 2 repo


	authHandler := handler.NewAuthHandler(authService)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)
	alumniHandler := handler.NewAlumniHandler(alumniService)
	pekerjaanHandler := handler.NewPekerjaanHandler(pekerjaanService)


	app := fiber.New()


	api := app.Group("/api")
	route.SetupAuthRoutes(api, authHandler)


	protected := api.Group("", middleware.AuthRequired())
	route.SetupMahasiswaRoutes(protected, mahasiswaHandler)
	route.SetupAlumniRoutes(protected, alumniHandler)
	route.SetupPekerjaanRoutes(protected, pekerjaanHandler)


	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server berjalan di port :%s", port)
	log.Fatal(app.Listen(":" + port))
}