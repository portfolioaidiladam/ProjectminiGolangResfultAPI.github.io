package repository

import (
	"aidiladam/belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

// CategoryRepository adalah kontrak repository untuk operasi data kategori.
type CategoryRepository interface {
	// Save menyimpan data kategori baru ke database.
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Update memperbarui data kategori yang sudah ada di database.
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Delete menghapus data kategori dari database.
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	// FindById mencari kategori berdasarkan ID.
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	// FindAll mengambil semua data kategori dari database.
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
