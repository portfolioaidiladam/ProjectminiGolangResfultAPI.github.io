package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CategoryController adalah kontrak controller untuk menangani HTTP request kategori.
type CategoryController interface {
	// Create menangani request untuk membuat kategori baru.
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// Update menangani request untuk memperbarui data kategori.
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// Delete menangani request untuk menghapus kategori.
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// FindById menangani request untuk mengambil data kategori berdasarkan ID.
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// FindAll menangani request untuk mengambil semua data kategori.
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
