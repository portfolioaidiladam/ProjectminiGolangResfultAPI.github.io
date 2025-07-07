package app

import (
	"aidiladam/belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

// NewDB membuat dan mengembalikan koneksi database MySQL yang sudah terkonfigurasi.
func NewDB() *sql.DB {
	// Membuka koneksi ke database MySQL
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/belajar_golang_restful_api")
	// Jika terjadi error saat membuka koneksi, panic
	helper.PanicIfError(err)

	// Mengatur jumlah maksimum koneksi idle
	db.SetMaxIdleConns(5)
	// Mengatur jumlah maksimum koneksi terbuka
	db.SetMaxOpenConns(20)
	// Mengatur waktu maksimum koneksi dapat digunakan
	db.SetConnMaxLifetime(60 * time.Minute)
	// Mengatur waktu maksimum koneksi idle
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
