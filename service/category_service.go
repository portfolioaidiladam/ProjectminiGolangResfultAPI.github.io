package service

import (
	"aidiladam/belajar-golang-restful-api/model/web"
	"context"
)

// CategoryService adalah kontrak service untuk operasi bisnis kategori.
type CategoryService interface {
	// Create membuat kategori baru berdasarkan request dan mengembalikan response kategori.
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	// Update memperbarui data kategori berdasarkan request dan mengembalikan response kategori.
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	// Delete menghapus kategori berdasarkan ID.
	Delete(ctx context.Context, categoryId int)
	// FindById mencari kategori berdasarkan ID dan mengembalikan response kategori.
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	// FindAll mengambil semua data kategori dan mengembalikannya dalam bentuk slice response.
	FindAll(ctx context.Context) []web.CategoryResponse
}
