package web

// CategoryCreateRequest adalah struct untuk request pembuatan kategori baru.
type CategoryCreateRequest struct {
	// Name adalah nama kategori yang akan dibuat
	// Validasi: required (wajib diisi), min=1 (minimal 1 karakter), max=100 (maksimal 100 karakter)
	Name string `validate:"required,min=1,max=100" json:"name"`
}
