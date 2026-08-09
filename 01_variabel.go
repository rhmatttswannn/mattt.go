package main

import "fmt"

func main() {
	// ============================================
	// TOPIK 1: VARIABEL DAN TIPE DATA
	// ============================================

	// Cara 1: Deklarasi dengan var
	var nama string = "Ali"
	var umur int = 20
	var tinggi float64 = 1.75
	var sudahKahwin bool = false

	fmt.Println("=== CARA 1: Guna 'var' ===")
	fmt.Println("Nama      :", nama)
	fmt.Println("Umur      :", umur)
	fmt.Println("Tinggi    :", tinggi, "m")
	fmt.Println("Sudah kahwin:", sudahKahwin)

	// Cara 2: Short declaration (lebih popular!)
	bandar := "Kuala Lumpur"
	gaji := 3500.50
	aktif := true

	fmt.Println("\n=== CARA 2: Short Declaration ':=' ===")
	fmt.Println("Bandar :", bandar)
	fmt.Println("Gaji   : RM", gaji)
	fmt.Println("Aktif  :", aktif)

	// Cara 3: Deklarasi banyak variabel sekaligus
	var (
		negara  = "Malaysia"
		benua   = "Asia"
		bahasa  = "Melayu"
	)

	fmt.Println("\n=== CARA 3: Deklarasi Banyak Variabel ===")
	fmt.Println("Negara :", negara)
	fmt.Println("Benua  :", benua)
	fmt.Println("Bahasa :", bahasa)

	// Konstanta - nilai yang TIDAK boleh berubah
	const PI = 3.14159
	const NAMA_APP = "BelajarGo"

	fmt.Println("\n=== KONSTANTA (const) ===")
	fmt.Println("Nilai PI  :", PI)
	fmt.Println("Nama App  :", NAMA_APP)

	// Tipe data lain
	var huruf byte = 'A'     // byte = uint8
	var besar int64 = 9999999999

	fmt.Println("\n=== TIPE DATA LAIN ===")
	fmt.Printf("Huruf : %c (nilai: %d)\n", huruf, huruf)
	fmt.Println("Nombor besar:", besar)

	// fmt.Printf untuk format output
	fmt.Println("\n=== FORMAT OUTPUT dengan Printf ===")
	fmt.Printf("Nama: %s, Umur: %d tahun, Tinggi: %.2f m\n", nama, umur, tinggi)
}
