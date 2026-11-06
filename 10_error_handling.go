package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrPembagiNol = errors.New("pembagi tidak boleh nol")

type ErrorMarkah struct {
	Nilai int
}

func (masalah *ErrorMarkah) Error() string {
	return fmt.Sprintf("markah %d di luar rentang 0 sampai 100", masalah.Nilai)
}

func bagi(angka, pembagi float64) (float64, error) {
	if pembagi == 0 {
		return 0, ErrPembagiNol
	}
	return angka / pembagi, nil
}

func bacaMarkah(teks string) (int, error) {
	markah, err := strconv.Atoi(teks)
	if err != nil {
		return 0, fmt.Errorf("membaca markah %q: %w", teks, err)
	}
	if markah < 0 || markah > 100 {
		return 0, &ErrorMarkah{Nilai: markah}
	}
	return markah, nil
}

func prosesMarkah(teks string) (int, error) {
	markah, err := bacaMarkah(teks)
	if err != nil {
		return 0, fmt.Errorf("memproses data pelajar: %w", err)
	}
	return markah, nil
}

func main() {
	fmt.Println("========== TOPIK 10: ERROR HANDLING ==========")
	fmt.Println("error adalah interface bawaan dengan method Error() string.")
	fmt.Println("Fungsi biasanya mengembalikan (hasil, error); nil berarti tidak ada error.")
	fmt.Println("Periksa error sebelum memakai hasil, dan jangan abaikan error dengan _.")

	fmt.Println("\n========== HASIL DAN ERROR ==========")
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Pembagian gagal:", err)
	} else {
		fmt.Println("10 / 2 =", hasil)
	}

	fmt.Println("\n========== ERRORS.NEW DAN ERRORS.IS ==========")
	hasil, err = bagi(10, 0)
	if err != nil {
		fmt.Println("Pembagian gagal:", err)
		fmt.Println("Penyebabnya pembagi nol:", errors.Is(err, ErrPembagiNol))
	} else {
		fmt.Println("Hasil:", hasil)
	}
	fmt.Println("ErrPembagiNol dibuat sekali agar bisa dikenali sebagai error yang sama.")
	fmt.Println("Jangan membandingkan pesan err.Error() untuk menentukan jenis error.")

	fmt.Println("\n========== WRAPPING DENGAN %w ==========")
	_, err = prosesMarkah("abc")
	if err != nil {
		fmt.Println("Pesan lengkap:", err)
		fmt.Println("Penyebabnya format angka:", errors.Is(err, strconv.ErrSyntax))
	}
	fmt.Println("fmt.Errorf dengan %w menambahkan konteks sambil mempertahankan error asal.")
	fmt.Println("errors.Is dapat menelusuri rantai error yang dibungkus.")

	fmt.Println("\n========== CUSTOM ERROR DAN ERRORS.AS ==========")
	_, err = prosesMarkah("120")
	if err != nil {
		fmt.Println("Pesan lengkap:", err)
		var masalah *ErrorMarkah
		if errors.As(err, &masalah) {
			fmt.Println("Nilai yang ditolak:", masalah.Nilai)
		}
	}
	fmt.Println("ErrorMarkah memenuhi interface error karena memiliki method Error() string.")
	fmt.Println("errors.As menemukan tipe error tertentu, termasuk di dalam wrapping.")

	fmt.Println("\n========== PRAKTIKAL: MEMPROSES BANYAK INPUT ==========")
	input := []string{"85", "abc", "120", "-5", "0", "100", "60"}
	jumlahValid := 0
	for _, teks := range input {
		markah, err := prosesMarkah(teks)
		if err != nil {
			var masalah *ErrorMarkah
			switch {
			case errors.Is(err, strconv.ErrSyntax):
				fmt.Printf("Input %q ditolak: masukkan bilangan bulat\n", teks)
			case errors.As(err, &masalah):
				fmt.Printf("Input %q ditolak: %d harus berada di rentang 0 sampai 100\n", teks, masalah.Nilai)
			default:
				fmt.Printf("Input %q gagal: %v\n", teks, err)
			}
			continue
		}
		jumlahValid++
		fmt.Printf("Input %q diterima, markah: %d\n", teks, markah)
	}
	fmt.Printf("Data valid: %d daripada %d input\n", jumlahValid, len(input))
	fmt.Println("Kegagalan satu input tidak menghentikan pemrosesan input berikutnya.")
	fmt.Println("Gunakan error untuk kegagalan yang wajar ditangani, bukan panic.")
}
