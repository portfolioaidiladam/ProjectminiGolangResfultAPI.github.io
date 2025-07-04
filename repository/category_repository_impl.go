package repository

import (
	"context"
	"database/sql"
	"errors"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

// CategoryRepositoryImpl adalah implementasi dari interface CategoryRepository untuk operasi data kategori.
type CategoryRepositoryImpl struct {
}

// NewCategoryRepository membuat instance baru dari CategoryRepositoryImpl.
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// Save menyimpan data kategori baru ke database dan mengembalikan kategori yang sudah memiliki ID.
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Menyusun query SQL untuk insert kategori baru
	SQL := "insert into category(name) values (?)"
	// Menjalankan query insert dengan parameter nama kategori
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	// Jika terjadi error saat eksekusi query, aplikasi akan panic
	helper.PanicIfError(err)

	// Mengambil ID hasil insert
	id, err := result.LastInsertId()
	// Jika terjadi error saat mengambil ID, aplikasi akan panic
	helper.PanicIfError(err)

	// Set ID hasil insert ke struct kategori
	category.Id = int(id)
	return category
}

// Update memperbarui data kategori yang sudah ada di database.
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Menyusun query SQL untuk update kategori berdasarkan id
	SQL := "update category set name = ? where id = ?"
	// Menjalankan query update dengan parameter nama dan id kategori
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	// Jika terjadi error saat eksekusi query, aplikasi akan panic
	helper.PanicIfError(err)

	return category
}

// Delete menghapus data kategori dari database berdasarkan ID.
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// Menyusun query SQL untuk menghapus kategori berdasarkan id
	SQL := "delete from category where id = ?"
	// Menjalankan query delete dengan parameter id kategori
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	// Jika terjadi error saat eksekusi query, aplikasi akan panic
	helper.PanicIfError(err)
}

// FindById mencari kategori berdasarkan ID dan mengembalikan data kategori jika ditemukan.
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// Menyusun query SQL untuk select kategori berdasarkan id
	SQL := "select id, name from category where id = ?"
	// Menjalankan query select dengan parameter id kategori
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	// Jika terjadi error saat eksekusi query, aplikasi akan panic
	helper.PanicIfError(err)
	// Pastikan rows ditutup setelah selesai
	defer rows.Close()

	category := domain.Category{}
	// Mengecek apakah data kategori ditemukan
	if rows.Next() {
		// Scan hasil query ke struct kategori
		err := rows.Scan(&category.Id, &category.Name)
		// Jika terjadi error saat scan, aplikasi akan panic
		helper.PanicIfError(err)
		return category, nil
	} else {
		// Jika tidak ditemukan, kembalikan error
		return category, errors.New("category is not found")
	}
}

// FindAll mengambil semua data kategori dari database dan mengembalikannya dalam bentuk slice.
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// Menyusun query SQL untuk select semua kategori
	SQL := "select id, name from category"
	// Menjalankan query select
	rows, err := tx.QueryContext(ctx, SQL)
	// Jika terjadi error saat eksekusi query, aplikasi akan panic
	helper.PanicIfError(err)
	// Pastikan rows ditutup setelah selesai
	defer rows.Close()

	var categories []domain.Category
	// Iterasi setiap baris hasil query
	for rows.Next() {
		category := domain.Category{}
		// Scan hasil query ke struct kategori
		err := rows.Scan(&category.Id, &category.Name)
		// Jika terjadi error saat scan, aplikasi akan panic
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
