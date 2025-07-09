// Package web mendefinisikan struktur response standar untuk API.
package web

// WebResponse adalah struktur standar untuk response API.
//
// Fields:
//
//	Code   - Kode status HTTP (misal: 200, 404, 500)
//	Status - Status response (misal: "OK", "NOT_FOUND", "ERROR")
//	Data   - Data yang dikirimkan dalam response, bisa berupa objek, array, atau pesan error
type WebResponse struct {
	Code   int         `json:"code"`   // Kode status HTTP
	Status string      `json:"status"` // Status response (OK, ERROR, dll)
	Data   interface{} `json:"data"`   // Payload data response
}
