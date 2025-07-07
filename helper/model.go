package helper

import (
	"aidiladam/belajar-golang-restful-api/model/domain"
	"aidiladam/belajar-golang-restful-api/model/web"
)

// ToCategoryResponse mengkonversi domain Category menjadi web CategoryResponse.
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,   // Mengambil ID dari domain category
		Name: category.Name, // Mengambil nama dari domain category
	}
}

// ToCategoryResponses mengkonversi slice domain Category menjadi slice web CategoryResponse.
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	// Iterasi setiap category dalam slice
	for _, category := range categories {
		// Konversi setiap category ke response dan tambahkan ke slice
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
