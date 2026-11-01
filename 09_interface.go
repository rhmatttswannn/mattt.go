package main

import (
	"fmt"
	"math"
)

type BangunDatar interface {
	Luas() float64
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

func (persegi PersegiPanjang) Luas() float64 {
	return persegi.Panjang * persegi.Lebar
}

type Lingkaran struct {
	JariJari float64
}

func (lingkaran Lingkaran) Luas() float64 {
	return math.Pi * lingkaran.JariJari * lingkaran.JariJari
}

func tampilkanLuas(bangun BangunDatar) {
	fmt.Printf("Jenis: %T | Luas: %.2f\n", bangun, bangun.Luas())
}

func totalLuas(daftar []BangunDatar) float64 {
	jumlah := 0.0
	for _, bangun := range daftar {
		jumlah += bangun.Luas()
	}
	return jumlah
}

func main() {
	fmt.Println("========== TOPIK 9: INTERFACE ==========")
	fmt.Println("Interface menentukan kontrak berupa kumpulan method.")
	fmt.Println("Tipe memenuhi interface jika memiliki semua method dengan signature yang sesuai.")
	fmt.Println("Tidak perlu menulis kata implements di Go.")

	fmt.Println("\n========== DUA STRUCT, SATU KONTRAK ==========")
	persegi := PersegiPanjang{Panjang: 5, Lebar: 3}
	lingkaran := Lingkaran{JariJari: 2}
	tampilkanLuas(persegi)
	tampilkanLuas(lingkaran)
	fmt.Println("Fungsi tampilkanLuas hanya membutuhkan method Luas(), bukan struct tertentu.")

	fmt.Println("\n========== VARIABEL INTERFACE ==========")
	var bangun BangunDatar = persegi
	fmt.Printf("Berisi %T, luas: %.2f\n", bangun, bangun.Luas())
	bangun = lingkaran
	fmt.Printf("Sekarang berisi %T, luas: %.2f\n", bangun, bangun.Luas())
	fmt.Println("Tipe variabel tetap BangunDatar; tipe konkret di dalamnya dapat berubah.")

	fmt.Println("\n========== SLICE BERISI INTERFACE ==========")
	daftar := []BangunDatar{
		persegi,
		lingkaran,
		PersegiPanjang{Panjang: 4, Lebar: 4},
	}
	for indeks, item := range daftar {
		fmt.Printf("Bangun ke-%d: ", indeks+1)
		tampilkanLuas(item)
	}
	fmt.Printf("Total luas: %.2f\n", totalLuas(daftar))

	fmt.Println("\n========== TYPE ASSERTION ==========")
	if hasil, cocok := bangun.(Lingkaran); cocok {
		fmt.Println("Jari-jari lingkaran:", hasil.JariJari)
	}
	_, cocok := bangun.(PersegiPanjang)
	fmt.Println("Isi bangun adalah PersegiPanjang:", cocok)
	fmt.Println("Pola nilai, cocok menghindari panic jika tipe tidak sesuai.")

	fmt.Println("\n========== TYPE SWITCH ==========")
	for _, item := range daftar {
		switch bentuk := item.(type) {
		case PersegiPanjang:
			fmt.Printf("Persegi panjang: panjang=%.1f, lebar=%.1f\n", bentuk.Panjang, bentuk.Lebar)
		case Lingkaran:
			fmt.Printf("Lingkaran: jari-jari=%.1f\n", bentuk.JariJari)
		default:
			fmt.Printf("Bangun lain: %T\n", bentuk)
		}
	}
	fmt.Println("Untuk menghitung luas, cukup panggil method tanpa memeriksa tipe.")
	fmt.Println("Assertion dan type switch berguna saat membutuhkan data khusus tipe konkret.")

	fmt.Println("\n========== POINTER DAN INTERFACE ==========")
	var melaluiPointer BangunDatar = &persegi
	fmt.Printf("Pointer %T juga memiliki method Luas(): %.2f\n", melaluiPointer, melaluiPointer.Luas())
	fmt.Println("Method dengan value receiver tersedia pada nilai dan pointer ke tipe itu.")
	fmt.Println("Jika Luas memakai pointer receiver, hanya pointer yang memenuhi kontrak ini.")

	fmt.Println("\n========== NIL INTERFACE ==========")
	var belumDiisi BangunDatar
	fmt.Println("Interface belum diisi bernilai nil:", belumDiisi == nil)
	fmt.Println("Jangan panggil method pada interface nil.")
	var pointerKosong *PersegiPanjang
	belumDiisi = pointerKosong
	fmt.Println("Interface berisi pointer nil bernilai nil:", belumDiisi == nil)
	fmt.Printf("Tipe konkret yang tersimpan: %T\n", belumDiisi)
	fmt.Println("Interface hanya nil jika tidak memiliki tipe konkret maupun nilai.")
	fmt.Println("Method Luas pada pointer nil ini juga tidak aman dipanggil.")
}
