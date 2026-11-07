package main

import "fmt"

func contohUrutanDefer() {
	fmt.Println("Mulai fungsi")
	defer fmt.Println("Defer pertama: dijalankan terakhir")
	defer fmt.Println("Defer kedua: dijalankan lebih dahulu")
	fmt.Println("Akhir badan fungsi")
}

func contohArgumenDefer() {
	nilai := 10
	defer fmt.Println("Argumen disimpan saat defer didaftarkan:", nilai)
	defer func() {
		fmt.Println("Closure membaca variabel saat dijalankan:", nilai)
	}()
	nilai = 20
	fmt.Println("Nilai sebelum fungsi selesai:", nilai)
}

func periksaPesanan(jumlah int) error {
	defer fmt.Println("Pemeriksaan pesanan selesai, termasuk jika return lebih awal")
	if jumlah <= 0 {
		return fmt.Errorf("jumlah pesanan harus positif, diterima %d", jumlah)
	}
	fmt.Println("Pesanan diterima:", jumlah)
	return nil
}

func simulasiPanic() {
	defer fmt.Println("Defer di simulasiPanic tetap dijalankan")
	fmt.Println("Sebelum panic di simulasiPanic")
	panic("contoh panic terkontrol untuk belajar")
}

func jalankanDenganRecover(tugas func()) (err error) {
	defer func() {
		if penyebab := recover(); penyebab != nil {
			err = fmt.Errorf("tugas mengalami panic: %v", penyebab)
		}
	}()
	tugas()
	fmt.Println("Tugas selesai normal tanpa panic")
	return nil
}

func main() {
	fmt.Println("========== TOPIK 11: DEFER, PANIC, DAN RECOVER ==========")
	fmt.Println("defer menunda pemanggilan fungsi sampai fungsi yang mendaftarkannya keluar.")
	fmt.Println("panic menghentikan alur normal dan menjalankan defer sambil menelusuri pemanggil.")
	fmt.Println("recover dapat menangkap panic dari fungsi deferred dalam goroutine yang sama.")

	fmt.Println("\n========== URUTAN DEFER: TERAKHIR MASUK, PERTAMA KELUAR ==========")
	contohUrutanDefer()
	fmt.Println("Kembali ke main setelah semua defer dalam contohUrutanDefer selesai")

	fmt.Println("\n========== ARGUMEN DEFER DAN CLOSURE ==========")
	contohArgumenDefer()
	fmt.Println("Argumen panggilan defer dievaluasi saat didaftarkan, bukan saat dieksekusi.")
	fmt.Println("Closure di atas membaca variabel nilai yang sama, yang sudah berubah menjadi 20.")

	fmt.Println("\n========== DEFER TETAP BERJALAN SAAT RETURN AWAL ==========")
	for _, jumlah := range []int{2, 0} {
		if err := periksaPesanan(jumlah); err != nil {
			fmt.Println("Pesanan ditolak:", err)
		}
	}
	fmt.Println("Defer terikat pada fungsi, bukan pada akhir blok if atau iterasi loop.")

	fmt.Println("\n========== RECOVER TANPA PANIC ==========")
	err := jalankanDenganRecover(func() {
		fmt.Println("Menjalankan tugas normal")
	})
	fmt.Println("Tugas normal menghasilkan error:", err)

	fmt.Println("\n========== PANIC TERKONTROL DAN RECOVER ==========")
	err = jalankanDenganRecover(simulasiPanic)
	if err != nil {
		fmt.Println("Ditangani oleh pemanggil:", err)
	}
	fmt.Println("Program berlanjut di main setelah jalankanDenganRecover selesai.")
	fmt.Println("Eksekusi tidak kembali ke baris setelah panic dalam simulasiPanic.")
	fmt.Println("Baris setelah tugas() pada jalur panic juga dilewati.")

	fmt.Println("\n========== KAPAN DIGUNAKAN? ==========")
	fmt.Println("Gunakan defer untuk pembersihan, misalnya menutup file setelah berhasil dibuka.")
	fmt.Println("Defer tidak berjalan jika program keluar melalui os.Exit.")
	fmt.Println("Panic yang tidak dipulihkan akan menghentikan program.")
	fmt.Println("Gunakan error untuk input tidak valid dan kegagalan yang wajar ditangani.")
	fmt.Println("Recover bukan pengganti pemeriksaan error dan tidak sebaiknya menyembunyikan bug.")
	fmt.Println("Wrapper contoh ini hanya untuk demonstrasi, bukan untuk membungkus semua fungsi.")
}
