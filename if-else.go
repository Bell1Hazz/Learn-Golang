package main

import "fmt"

func main() {
	produk := "Laptop Gaming"
	harga := 15000000
	tersedia := true // Kita ubah jadi true untuk latihan ini

	fmt.Println("Nama Produk:", produk)
	fmt.Println("Harga: Rp", harga)

	// 1. If-Else Sederhana (Cek Ketersediaan)
	if tersedia == true {
		fmt.Println("Status: Silakan dipesan, stok tersedia!")
	} else {
		fmt.Println("Status: Maaf, stok habis.")
	}

	// 2. If-Else Bertingkat / Else If (Cek Tingkat Harga)
	if harga > 20000000 {
		fmt.Println("Kategori: Laptop Sultan / Premium")
	} else if harga > 10000000 {
		fmt.Println("Kategori: Laptop Mid-Range / Menengah")
	} else {
		fmt.Println("Kategori: Laptop Entry-Level / Murah")
	}
}
