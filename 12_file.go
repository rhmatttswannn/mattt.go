package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func tambahCatatan(lokasi, teks string) (err error) {
	file, err := os.OpenFile(lokasi, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("membuka file untuk append: %w", err)
	}
	defer func() {
		if masalah := file.Close(); masalah != nil && err == nil {
			err = fmt.Errorf("menutup file catatan: %w", masalah)
		}
	}()
	if _, err := file.WriteString(teks); err != nil {
		return fmt.Errorf("menambah catatan: %w", err)
	}
	return nil
}

func bacaBaris(lokasi string) (baris []string, err error) {
	file, err := os.Open(lokasi)
	if err != nil {
		return nil, fmt.Errorf("membuka file untuk dibaca: %w", err)
	}
	defer func() {
		if masalah := file.Close(); masalah != nil && err == nil {
			err = fmt.Errorf("menutup file bacaan: %w", masalah)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		baris = append(baris, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("membaca baris: %w", err)
	}
	return baris, nil
}

func belajarFile() (err error) {
	fmt.Println("========== TOPIK 12: BACA DAN TULIS FILE ==========")
	folder, err := os.MkdirTemp("", "belajar-go-file-*")
	if err != nil {
		return fmt.Errorf("membuat folder sementara: %w", err)
	}
	defer func() {
		if masalah := os.RemoveAll(folder); masalah != nil {
			if err == nil {
				err = fmt.Errorf("membersihkan folder sementara: %w", masalah)
			}
			return
		}
		fmt.Println("Folder sementara sudah dibersihkan.")
	}()
	lokasi := filepath.Join(folder, "catatan.txt")
	fmt.Println("Lokasi latihan:", lokasi)
	fmt.Println("filepath.Join menyusun path sesuai sistem operasi.")
	fmt.Println("Semua file latihan berada dalam folder sementara milik program ini.")

	fmt.Println("\n========== MENULIS FILE ==========")
	isi := "Belajar Go\nMemahami error dan defer\n"
	if err := os.WriteFile(lokasi, []byte(isi), 0600); err != nil {
		return fmt.Errorf("menulis catatan: %w", err)
	}
	fmt.Println("Dua baris awal berhasil ditulis.")
	fmt.Println("WriteFile membuat file baru atau menimpa seluruh isi file yang sudah ada.")
	fmt.Println("0600 meminta izin baca/tulis untuk pemilik pada sistem Unix; Windows berbeda.")

	fmt.Println("\n========== MEMBACA SELURUH ISI ==========")
	data, err := os.ReadFile(lokasi)
	if err != nil {
		return fmt.Errorf("membaca catatan: %w", err)
	}
	fmt.Print(string(data))
	fmt.Println("ReadFile menghasilkan []byte; string(data) mengubahnya menjadi teks.")
	fmt.Println("WriteFile dan ReadFile mengelola buka/tutup file secara internal.")

	fmt.Println("\n========== MENAMBAHKAN ISI TANPA MENIMPA ==========")
	if err := tambahCatatan(lokasi, "Praktik membaca dan menulis file\n"); err != nil {
		return err
	}
	fmt.Println("O_APPEND menulis di akhir file, O_WRONLY membuka untuk menulis.")
	fmt.Println("Tanpa O_CREATE, OpenFile di contoh ini gagal jika file belum ada.")

	fmt.Println("\n========== MEMBACA PER BARIS ==========")
	baris, err := bacaBaris(lokasi)
	if err != nil {
		return err
	}
	for indeks, teks := range baris {
		fmt.Printf("Baris %d: %s\n", indeks+1, teks)
	}
	fmt.Println("Jumlah baris:", len(baris))
	fmt.Println("Scanner.Text tidak menyertakan pemisah baris.")
	fmt.Println("Selalu periksa scanner.Err setelah loop Scan selesai.")
	fmt.Println("Scanner memiliki batas ukuran token bawaan sekitar 64 KiB, bukan untuk baris tak terbatas.")
	fmt.Println("Open dan OpenFile perlu ditutup; daftarkan defer setelah pembukaan berhasil.")

	fmt.Println("\n========== FILE TIDAK DITEMUKAN ==========")
	_, err = bacaBaris(filepath.Join(folder, "belum_ada.txt"))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		fmt.Println("File belum ada, dikenali melalui errors.Is(err, os.ErrNotExist).")
	}
	fmt.Println("Latihan selesai; defer akan membersihkan folder sementara.")
	return nil
}

func main() {
	if err := belajarFile(); err != nil {
		fmt.Println("Latihan file gagal:", err)
	}
}
