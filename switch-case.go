package main

import "fmt"

func main() {
	// Kita buat variabel baru untuk menentukan metode pembayaran
	metodePembayaran := "Transfer Bank"

	fmt.Println("Metode Pembayaran Pilihan:", metodePembayaran)

	// Menggunakan Switch-Case
	switch metodePembayaran {
	case "Transfer Bank":
		fmt.Println("Petunjuk: Silakan transfer ke rekening BCA 123456.")
	case "Gopay":
		fmt.Println("Petunjuk: Silakan scan QRIS yang muncul di layar.")
	case "OVO":
		fmt.Println("Petunjuk: Silakan cek notifikasi di aplikasi OVO kamu.")
	default:
		fmt.Println("Petunjuk: Metode pembayaran tidak dikenali.")
	}
}

//Untuk melakukan testing tinggal ganti value metodePembayaran agar sesuai Case
