package main

import "fmt"

func main() {
	produk := "Laptop Gaming"
	harga := 15000000
	tersedia := true
	memberToko := true
	metodePembayaran := "Transfer"

	fmt.Println("--- Transaksi:", produk, "---")

	// 1. Contoh Penggunaan Operator && (AND)
	// Syarat dapat bonus: Stok harus ada DAN harganya harus di atas 10 juta
	if tersedia == true && harga > 10000000 {
		fmt.Println("Selamat! Anda mendapatkan Bonus Mouse Gaming.")
	} else {
		fmt.Println("Maaf, Anda belum memenuhi syarat mendapat bonus.")
	}

	// 2. Contoh Penggunaan Operator || (OR)
	// Syarat dapat diskon: Harus jadi member toko ATAU bayar pakai Gopay
	if memberToko == true || metodePembayaran == "Gopay" {
		fmt.Println("Diskon: Anda mendapatkan potongan harga Rp 500.000!")
	} else {
		fmt.Println("Diskon: Harga normal (tidak ada diskon).")
	}

	// 3. Contoh Penggunaan Operator ! (NOT)
	// Membalikkan logika: jika "TIDAK tersedia"
	if !tersedia {
		fmt.Println("Pemberitahuan: Produk ini sedang tidak bisa dibeli.")
	} else {
		fmt.Println("Pemberitahuan: Produk siap diangkut!")
	}
}
