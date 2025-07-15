package web

// CategoryResponse adalah struct untuk response data kategori.
type CategoryResponse struct {
	// Id adalah identitas unik kategori
	Id int `json:"id"`
	// Name adalah nama kategori
	Name string `json:"name"`
}
