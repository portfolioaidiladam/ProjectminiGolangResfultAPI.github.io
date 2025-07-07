package service

import (
	"aidiladam/belajar-golang-restful-api/exception"
	"aidiladam/belajar-golang-restful-api/helper"
	"aidiladam/belajar-golang-restful-api/model/domain"
	"aidiladam/belajar-golang-restful-api/model/web"
	"aidiladam/belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

// CategoryServiceImpl adalah implementasi dari interface CategoryService untuk operasi bisnis kategori.
type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository // Repository untuk operasi data kategori
	DB                 *sql.DB                      // Koneksi database
	Validate           *validator.Validate           // Validator untuk validasi request
}

// NewCategoryService membuat instance baru dari CategoryServiceImpl.
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Create membuat kategori baru berdasarkan request dan mengembalikan response kategori.
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// Validasi request menggunakan validator
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// Commit atau rollback transaksi secara otomatis
	defer helper.CommitOrRollback(tx)

	// Membuat entity kategori dari request
	category := domain.Category{
		Name: request.Name,
	}

	// Simpan kategori ke database
	category = service.CategoryRepository.Save(ctx, tx, category)

	// Konversi entity kategori ke response
	return helper.ToCategoryResponse(category)
}

// Update memperbarui data kategori berdasarkan request dan mengembalikan response kategori.
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// Validasi request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cari kategori berdasarkan ID
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		// Jika tidak ditemukan, panic dengan error not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update nama kategori
	category.Name = request.Name

	// Simpan perubahan ke database
	category = service.CategoryRepository.Update(ctx, tx, category)

	// Konversi entity kategori ke response
	return helper.ToCategoryResponse(category)
}

// Delete menghapus kategori berdasarkan ID.
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cari kategori berdasarkan ID
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		// Jika tidak ditemukan, panic dengan error not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Hapus kategori dari database
	service.CategoryRepository.Delete(ctx, tx, category)
}

// FindById mencari kategori berdasarkan ID dan mengembalikan response kategori.
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cari kategori berdasarkan ID
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		// Jika tidak ditemukan, panic dengan error not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Konversi entity kategori ke response
	return helper.ToCategoryResponse(category)
}

// FindAll mengambil semua data kategori dan mengembalikannya dalam bentuk slice response.
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Ambil semua kategori dari database
	categories := service.CategoryRepository.FindAll(ctx, tx)

	// Konversi slice entity kategori ke slice response
	return helper.ToCategoryResponses(categories)
}
