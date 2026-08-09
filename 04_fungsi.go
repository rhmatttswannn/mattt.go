package main

import "fmt"

// ============================================
// TOPIK 4: FUNGSI (FUNCTION)
// ============================================

// Fungsi mudah - tiada parameter, tiada return value
func sapa() {
	fmt.Println("Helo! Selamat datang ke dunia Go!")
}

// Fungsi dengan parameter
func sapaNama(nama string) {
	fmt.Printf("Helo, %s! Apa khabar?\n", nama)
}

// Fungsi dengan parameter dan return value
func tambah(a int, b int) int {
	return a + b
}

// Ringkas: jika tipe sama, boleh tulis sekali
func tolak(a, b int) int {
	return a - b
}

// Fungsi dengan BANYAK return value (keunikan Go!)
func bahagi(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "ERROR: Tidak boleh bahagi dengan 0!"
	}
	hasil := a / b
	return hasil, "OK"
}

// Named return values
func kiraSegi(panjang, lebar float64) (luas, perimeter float64) {
	luas = panjang * lebar
	perimeter = 2 * (panjang + lebar)
	return // return tanpa nilai - akan return luas & perimeter
}

// Fungsi rekursif (fungsi panggil dirinya sendiri)
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

// Fungsi sebagai parameter (higher-order function)
func operasi(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func main() {
	fmt.Println("========== FUNGSI MUDAH ==========")
	sapa()
	sapaNama("Ali")
	sapaNama("Siti")

	fmt.Println("\n========== FUNGSI DENGAN RETURN ==========")
	jumlah := tambah(10, 5)
	beza := tolak(10, 5)
	fmt.Println("10 + 5 =", jumlah)
	fmt.Println("10 - 5 =", beza)

	fmt.Println("\n========== MULTIPLE RETURN VALUES ==========")
	hasil, status := bahagi(10, 3)
	fmt.Printf("10 / 3 = %.4f | Status: %s\n", hasil, status)

	hasil2, status2 := bahagi(10, 0)
	fmt.Printf("10 / 0 = %.4f | Status: %s\n", hasil2, status2)

	fmt.Println("\n========== NAMED RETURN ==========")
	luas, perimeter := kiraSegi(5, 3)
	fmt.Printf("Segi empat 5x3 → Luas: %.1f | Perimeter: %.1f\n", luas, perimeter)

	fmt.Println("\n========== FUNGSI REKURSIF ==========")
	for i := 0; i <= 7; i++ {
		fmt.Printf("%d! = %d\n", i, faktorial(i))
	}

	fmt.Println("\n========== FUNGSI SEBAGAI PARAMETER ==========")
	// Hantar fungsi tambah sebagai argumen
	hasilOp := operasi(8, 4, tambah)
	fmt.Println("operasi(8, 4, tambah) =", hasilOp)

	hasilOp2 := operasi(8, 4, tolak)
	fmt.Println("operasi(8, 4, tolak)  =", hasilOp2)

	// Anonymous function (fungsi tanpa nama)
	fmt.Println("\n========== ANONYMOUS FUNCTION ==========")
	kali := func(a, b int) int {
		return a * b
	}
	fmt.Println("kali(6, 7) =", kali(6, 7))

	// Terus panggil tanpa simpan ke variabel
	hasil3 := func(a, b int) int {
		return a * a + b * b
	}(3, 4) // terus bagi nilai
	fmt.Println("3² + 4² =", hasil3)
}
