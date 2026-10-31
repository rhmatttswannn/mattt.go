package main

import "fmt"

func main() {
	// ============================================
	// TOPIK 5: ARRAY & SLICE
	// ============================================
	// Array  = saiz TETAP, tidak boleh berubah
	// Slice  = saiz DINAMIK, boleh tambah/kurang

	// ==========================================
	// BAHAGIAN A: ARRAY
	// ==========================================
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║          A R R A Y           ║")
	fmt.Println("╚══════════════════════════════╝")

	// Deklarasi array - saiz mesti ditetapkan
	var markah [5]int
	markah[0] = 80
	markah[1] = 95
	markah[2] = 70
	markah[3] = 88
	markah[4] = 60

	fmt.Println("\n--- Array Markah ---")
	fmt.Println("Array penuh :", markah)
	fmt.Println("Elemen ke-0 :", markah[0])
	fmt.Println("Elemen ke-4 :", markah[4])
	fmt.Println("Saiz array  :", len(markah))

	// Deklarasi dan isi terus
	buah := [4]string{"Epal", "Mangga", "Durian", "Rambutan"}
	fmt.Println("\n--- Array Buah ---")
	fmt.Println("Semua buah:", buah)

	// Loop through array
	fmt.Println("\n--- Loop Array dengan index ---")
	for i := 0; i < len(buah); i++ {
		fmt.Printf("buah[%d] = %s\n", i, buah[i])
	}

	// Range - cara lebih Go untuk loop array
	fmt.Println("\n--- Loop Array dengan 'range' ---")
	for index, nilai := range markah {
		fmt.Printf("markah[%d] = %d\n", index, nilai)
	}

	// Guna _ jika tak perlukan index
	fmt.Println("\n--- Hanya nilai sahaja (guna _) ---")
	jumlah := 0
	for _, nilai := range markah {
		jumlah += nilai
	}
	purata := float64(jumlah) / float64(len(markah))
	fmt.Printf("Jumlah: %d | Purata: %.1f\n", jumlah, purata)

	// Array 2 dimensi (macam jadual/grid)
	fmt.Println("\n--- Array 2D (Jadual 3x3) ---")
	grid := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, baris := range grid {
		for _, sel := range baris {
			fmt.Printf("%d ", sel)
		}
		fmt.Println()
	}

	// ==========================================
	// BAHAGIAN B: SLICE
	// ==========================================
	fmt.Println("\n╔══════════════════════════════╗")
	fmt.Println("║          S L I C E           ║")
	fmt.Println("╚══════════════════════════════╝")

	// Beza slice vs array: tiada saiz dinyatakan
	nama := []string{"Ali", "Siti", "Kumar", "Mei"}
	fmt.Println("\n--- Slice Nama ---")
	fmt.Println("Slice   :", nama)
	fmt.Println("Panjang :", len(nama)) // bilangan elemen semasa
	fmt.Println("Kapasiti:", cap(nama)) // kapasiti dalaman

	// Tambah elemen dengan append
	fmt.Println("\n--- Guna append() ---")
	nama = append(nama, "Ravi")
	nama = append(nama, "Nurul", "Hafiz") // tambah banyak sekaligus
	fmt.Println("Selepas append:", nama)
	fmt.Println("Panjang baru  :", len(nama))

	// Slicing - ambil sebahagian dari slice
	fmt.Println("\n--- Slicing (ambil sebahagian) ---")
	nombor := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println("Asal        :", nombor)
	fmt.Println("nombor[1:4] :", nombor[1:4]) // index 1 hingga 3
	fmt.Println("nombor[:3]  :", nombor[:3])  // dari awal hingga index 2
	fmt.Println("nombor[4:]  :", nombor[4:])  // dari index 4 hingga akhir
	fmt.Println("nombor[:]   :", nombor[:])   // semua

	// make() - buat slice dengan saiz & kapasiti tertentu
	fmt.Println("\n--- Guna make() ---")
	s := make([]int, 3, 5) // panjang=3, kapasiti=5
	fmt.Printf("make([]int, 3, 5) → %v | len=%d | cap=%d\n", s, len(s), cap(s))

	// Copy slice
	fmt.Println("\n--- Copy Slice ---")
	asal := []int{1, 2, 3, 4, 5}
	salinan := make([]int, len(asal))
	copy(salinan, asal)
	salinan[0] = 999 // ubah salinan, asal tidak terjejas
	fmt.Println("Asal   :", asal)
	fmt.Println("Salinan:", salinan)

	// Padam elemen dari slice
	fmt.Println("\n--- Padam Elemen dari Slice ---")
	buah2 := []string{"Epal", "Mangga", "Durian", "Rambutan", "Betik"}
	fmt.Println("Sebelum padam:", buah2)

	// Padam index ke-2 (Durian)
	padamIndex := 2
	buah2 = append(buah2[:padamIndex], buah2[padamIndex+1:]...)
	fmt.Printf("Selepas padam index %d: %v\n", padamIndex, buah2)

	// ==========================================
	// CONTOH PRAKTIKAL
	// ==========================================
	fmt.Println("\n╔══════════════════════════════╗")
	fmt.Println("║    CONTOH PRAKTIKAL          ║")
	fmt.Println("╚══════════════════════════════╝")

	// Cari nilai minimum & maksimum
	skor := []int{45, 78, 23, 90, 56, 12, 88, 67}
	min := skor[0]
	maks := skor[0]

	for _, nilai := range skor {
		if nilai < min {
			min = nilai
		}
		if nilai > maks {
			maks = nilai
		}
	}
	fmt.Println("\nSenarai skor:", skor)
	fmt.Println("Nilai minimum:", min)
	fmt.Println("Nilai maksimum:", maks)

	// Tapis (filter) elemen yang memenuhi syarat
	fmt.Println("\n--- Filter: hanya skor >= 60 (Lulus) ---")
	lulus := []int{}
	gagal := []int{}

	for _, nilai := range skor {
		if nilai >= 60 {
			lulus = append(lulus, nilai)
		} else {
			gagal = append(gagal, nilai)
		}
	}
	fmt.Println("Lulus:", lulus)
	fmt.Println("Gagal:", gagal)

	// Balikkan (reverse) slice
	fmt.Println("\n--- Balikkan Slice ---")
	perkataan := []string{"Go", "adalah", "bahasa", "yang", "pantas"}
	fmt.Println("Asal   :", perkataan)

	for i, j := 0, len(perkataan)-1; i < j; i, j = i+1, j-1 {
		perkataan[i], perkataan[j] = perkataan[j], perkataan[i]
	}
	fmt.Println("Balik  :", perkataan)
}
