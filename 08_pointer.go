package main

import "fmt"

func ubahSalinan(nilai int) {
	nilai = 100
	fmt.Println("Nilai di dalam ubahSalinan:", nilai)
}

func ubahMelaluiPointer(target *int, nilaiBaru int) bool {
	if target == nil {
		return false
	}
	*target = nilaiBaru
	return true
}

func tukarNilai(pertama, kedua *int) bool {
	if pertama == nil || kedua == nil {
		return false
	}
	*pertama, *kedua = *kedua, *pertama
	return true
}

func main() {
	fmt.Println("========== TOPIK 8: POINTER ==========")
	fmt.Println("Pointer menyimpan alamat suatu nilai.")
	fmt.Println("& mengambil alamat, * membaca atau mengubah nilai pada alamat itu.")

	fmt.Println("\n========== ALAMAT DAN NILAI ==========")
	angka := 10
	var penunjuk *int = &angka
	fmt.Println("Nilai angka:", angka)
	fmt.Printf("Alamat angka: %p\n", &angka)
	fmt.Printf("Alamat yang disimpan penunjuk: %p\n", penunjuk)
	fmt.Printf("Jenis penunjuk: %T\n", penunjuk)
	fmt.Println("Nilai melalui *penunjuk:", *penunjuk)
	fmt.Println("Alamat boleh berbeza setiap kali program dijalankan.")

	fmt.Println("\n========== MENGUBAH MELALUI POINTER ==========")
	*penunjuk = 25
	fmt.Println("Selepas *penunjuk = 25, angka:", angka)
	angka = 30
	fmt.Println("Selepas angka = 30, *penunjuk:", *penunjuk)

	fmt.Println("\n========== NILAI SEBAGAI ARGUMEN ==========")
	ubahSalinan(angka)
	fmt.Println("Angka asal selepas ubahSalinan:", angka)

	fmt.Println("\n========== POINTER SEBAGAI ARGUMEN ==========")
	berjaya := ubahMelaluiPointer(&angka, 100)
	fmt.Println("Perubahan diterima:", berjaya)
	fmt.Println("Angka asal selepas ubahMelaluiPointer:", angka)
	fmt.Println("Go sentiasa menyalin argumen, termasuk nilai pointer.")
	fmt.Println("Salinan pointer masih menunjuk data yang sama, jadi *target boleh mengubahnya.")

	fmt.Println("\n========== MENYALIN DAN MENGALIHKAN POINTER ==========")
	penunjukKedua := penunjuk
	*penunjukKedua = 150
	fmt.Println("Ubah melalui salinan pointer, angka menjadi:", angka)
	angkaLain := 200
	penunjukKedua = &angkaLain
	*penunjukKedua = 250
	fmt.Println("Selepas pointer kedua dialihkan, angka asal:", angka)
	fmt.Println("Angka lain:", angkaLain)
	fmt.Println("Menukar alamat dalam satu pointer tidak mengalihkan pointer yang lain.")

	fmt.Println("\n========== NIL POINTER ==========")
	var kosong *int
	fmt.Println("Pointer belum menunjuk nilai:", kosong == nil)
	fmt.Println("Membaca atau menulis *kosong akan menyebabkan panic.")
	berjaya = ubahMelaluiPointer(kosong, 50)
	fmt.Println("Perubahan melalui nil diterima:", berjaya)
	fmt.Println("Fungsi tadi memeriksa nil sebelum menggunakan *target.")

	fmt.Println("\n========== MEMBUAT POINTER DENGAN NEW ==========")
	alokasi := new(int)
	fmt.Println("Nilai awal dari new(int):", *alokasi)
	*alokasi = 75
	fmt.Println("Selepas diubah:", *alokasi)
	fmt.Println("new(int) menghasilkan *int yang menunjuk nilai awal 0, bukan nil.")

	fmt.Println("\n========== PRAKTIKAL: MENUKAR DUA NILAI ==========")
	nilaiPertama, nilaiKedua := 10, 20
	fmt.Printf("Sebelum: pertama=%d, kedua=%d\n", nilaiPertama, nilaiKedua)
	berjaya = tukarNilai(&nilaiPertama, &nilaiKedua)
	fmt.Println("Pertukaran diterima:", berjaya)
	fmt.Printf("Selepas: pertama=%d, kedua=%d\n", nilaiPertama, nilaiKedua)
	berjaya = tukarNilai(&nilaiPertama, nil)
	fmt.Println("Pertukaran dengan nil diterima:", berjaya)
	fmt.Printf("Nilai kekal: pertama=%d, kedua=%d\n", nilaiPertama, nilaiKedua)
}
