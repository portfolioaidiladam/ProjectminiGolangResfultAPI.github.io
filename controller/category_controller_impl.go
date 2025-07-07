package controller

import (
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// CategoryControllerImpl adalah implementasi dari interface CategoryController untuk menangani HTTP request kategori.
type CategoryControllerImpl struct {
	CategoryService service.CategoryService // Service untuk operasi bisnis kategori
}

// NewCategoryController membuat instance baru dari CategoryControllerImpl.
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// Create menangani request untuk membuat kategori baru.
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Membaca request body dan mapping ke struct CategoryCreateRequest
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	// Memanggil service untuk membuat kategori
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	// Membungkus response ke dalam WebResponse
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Menulis response ke response writer
	helper.WriteToResponseBody(writer, webResponse)
}

// Update menangani request untuk memperbarui data kategori.
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Membaca request body dan mapping ke struct CategoryUpdateRequest
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// Mengambil parameter categoryId dari URL
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// Set ID ke struct request
	categoryUpdateRequest.Id = id

	// Memanggil service untuk update kategori
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	// Membungkus response ke dalam WebResponse
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Menulis response ke response writer
	helper.WriteToResponseBody(writer, webResponse)
}

// Delete menangani request untuk menghapus kategori.
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Mengambil parameter categoryId dari URL
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// Memanggil service untuk menghapus kategori
	controller.CategoryService.Delete(request.Context(), id)
	// Membungkus response ke dalam WebResponse
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// Menulis response ke response writer
	helper.WriteToResponseBody(writer, webResponse)
}

// FindById menangani request untuk mengambil data kategori berdasarkan ID.
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Mengambil parameter categoryId dari URL
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// Memanggil service untuk mengambil data kategori
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	// Membungkus response ke dalam WebResponse
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Menulis response ke response writer
	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll menangani request untuk mengambil semua data kategori.
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Memanggil service untuk mengambil semua data kategori
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	// Membungkus response ke dalam WebResponse
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	// Menulis response ke response writer
	helper.WriteToResponseBody(writer, webResponse)
}
