package web

// CategoryUpdateRequest adalah struct untuk request pembaruan data kategori.
type CategoryUpdateRequest struct {
	// Id adalah identitas unik kategori yang akan diperbarui.
	// Validasi: required (wajib diisi)
	Id   int    `validate:"required"`
	// Name adalah nama kategori yang akan diperbarui.
	// Validasi: required (wajib diisi), min=1 (minimal 1 karakter), max=200 (maksimal 200 karakter)
	Name string `validate:"required,max=200,min=1" json:"name"`
}
