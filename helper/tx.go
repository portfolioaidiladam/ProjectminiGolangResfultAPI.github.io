package helper

import "database/sql"

// CommitOrRollback menangani commit atau rollback transaksi database secara otomatis.
// Fungsi ini menggunakan defer untuk memastikan transaksi selalu diakhiri dengan benar.
func CommitOrRollback(tx *sql.Tx) {
	// Cek apakah ada panic yang terjadi
	err := recover()
	if err != nil {
		// Jika ada panic, lakukan rollback transaksi
		errorRollback := tx.Rollback()
		// Jika rollback gagal, panic dengan error rollback
		PanicIfError(errorRollback)
		// Panic kembali dengan error asli
		panic(err)
	} else {
		// Jika tidak ada panic, lakukan commit transaksi
		errorCommit := tx.Commit()
		// Jika commit gagal, panic dengan error commit
		PanicIfError(errorCommit)
	}
}
