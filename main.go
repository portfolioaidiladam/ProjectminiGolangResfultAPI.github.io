package main

import (
	"aidiladam/belajar-golang-restful-api/app"
	"aidiladam/belajar-golang-restful-api/controller"
	"aidiladam/belajar-golang-restful-api/helper"

	"aidiladam/belajar-golang-restful-api/middleware"
	"aidiladam/belajar-golang-restful-api/repository"
	"aidiladam/belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Membuka koneksi ke database
	db := app.NewDB()
	// Membuat instance validator untuk validasi request
	validate := validator.New()
	// Membuat repository kategori
	categoryRepository := repository.NewCategoryRepository()
	// Membuat service kategori dengan dependency repository, db, dan validator
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// Membuat controller kategori dengan dependency service
	categoryController := controller.NewCategoryController(categoryService)
	// Membuat router dan mendaftarkan controller kategori
	router := app.NewRouter(categoryController)

	// Membuat server HTTP dengan middleware autentikasi
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// Menjalankan server HTTP
	err := server.ListenAndServe()
	// Jika terjadi error saat menjalankan server, panic
	helper.PanicIfError(err)
}
