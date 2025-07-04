package helper

// PanicIfError akan mem-panic aplikasi jika error tidak nil.
//
// Digunakan untuk memudahkan error handling pada operasi yang harus gagal total jika terjadi error.
func PanicIfError(err error) {
	if err != nil {
		panic(err) // Panic jika error ditemukan
	}
}
