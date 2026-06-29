package main

import "fmt"

func main() {
	// Membuat variabel tanpa nilai awal
	var namaToko string
	var totalStok int
	var diskonTersedia bool

	// Mencetak nilai bawaan (Zero Value)
	fmt.Println("Zero Value String  :", namaToko)
	fmt.Println("Zero Value Integer :", totalStok)
	fmt.Println("Zero Value Boolean :", diskonTersedia)

	fmt.Println("--------------------------------")

	// Mengisi nilai variabel setelahnya
	namaToko = "Tech Store"
	totalStok = 50
	diskonTersedia = true

	fmt.Println("Nama Toko Baru     :", namaToko)
	fmt.Println("Total Stok Baru    :", totalStok)
	fmt.Println("Status Diskon Baru :", diskonTersedia)
}
