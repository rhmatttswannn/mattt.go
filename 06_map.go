package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("========== TOPIK 6: MAP ==========")
	fmt.Println("Map menyimpan pasangan kunci dan nilai.")
	fmt.Println("Slice memakai indeks angka, map boleh memakai kunci seperti nama.")

	fmt.Println("\n========== MEMBUAT DAN MEMBACA MAP ==========")
	markah := map[string]int{
		"Ali":  80,
		"Siti": 95,
		"Budi": 70,
	}
	fmt.Println("Markah Ali:", markah["Ali"])
	fmt.Println("Jumlah pelajar:", len(markah))

	fmt.Println("\n========== MENAMBAH DAN MENGUBAH NILAI ==========")
	markah["Dewi"] = 88
	markah["Ali"] = 85
	fmt.Println("Markah Dewi:", markah["Dewi"])
	fmt.Println("Markah Ali selepas diubah:", markah["Ali"])

	fmt.Println("\n========== MEMERIKSA KUNCI ==========")
	markah["Rina"] = 0
	for _, nama := range []string{"Rina", "Andi"} {
		if nilai, ada := markah[nama]; ada {
			fmt.Printf("%s terdaftar, markah: %d\n", nama, nilai)
		} else {
			fmt.Printf("%s tidak terdaftar\n", nama)
		}
	}
	fmt.Println("Kunci yang tidak ada menghasilkan nilai sifar bagi jenis nilainya.")
	fmt.Println("Gunakan nilai, ada := map[kunci] untuk membezakan 0 dengan data yang tiada.")

	fmt.Println("\n========== MENGHAPUS DATA ==========")
	delete(markah, "Budi")
	delete(markah, "TidakAda")
	_, budiAda := markah["Budi"]
	fmt.Println("Budi masih terdaftar:", budiAda)
	fmt.Println("Jumlah pelajar selepas dihapus:", len(markah))

	fmt.Println("\n========== LOOP MAP ==========")
	fmt.Println("Urutan range pada map tidak dijamin tetap.")
	for nama, nilai := range markah {
		fmt.Printf("%s: %d\n", nama, nilai)
	}

	fmt.Println("\n========== PAPARAN TERURUT ==========")
	senaraiNama := make([]string, 0, len(markah))
	for nama := range markah {
		senaraiNama = append(senaraiNama, nama)
	}
	sort.Strings(senaraiNama)
	for _, nama := range senaraiNama {
		fmt.Printf("%s: %d\n", nama, markah[nama])
	}

	fmt.Println("\n========== MAKE DAN NIL MAP ==========")
	stok := make(map[string]int)
	stok["Buku"] = 10
	stok["Pensel"] = 5
	fmt.Println("Stok buku:", stok["Buku"])

	var belumSiap map[string]int
	fmt.Println("Map belum dibuat bernilai nil:", belumSiap == nil)
	fmt.Println("Membaca nil map dibenarkan:", belumSiap["Buku"])
	fmt.Println("Menulis ke nil map menyebabkan panic; buat map dengan make dahulu.")
	belumSiap = make(map[string]int)
	belumSiap["Buku"] = 3
	fmt.Println("Selepas make, stok buku:", belumSiap["Buku"])

	fmt.Println("\n========== MAP BERKONGSI DATA ==========")
	stokBersama := stok
	stokBersama["Buku"] = 20
	fmt.Println("Selepas ubah stokBersama, stok asal turut berubah:", stok["Buku"])

	salinanStok := make(map[string]int, len(stok))
	for barang, jumlah := range stok {
		salinanStok[barang] = jumlah
	}
	salinanStok["Buku"] = 99
	fmt.Println("Stok asal:", stok["Buku"])
	fmt.Println("Salinan berasingan:", salinanStok["Buku"])

	fmt.Println("\n========== PRAKTIKAL: FREKUENSI KATA ==========")
	ayat := "go mudah go seronok belajar go"
	frekuensi := make(map[string]int)
	for _, kata := range strings.Fields(ayat) {
		frekuensi[kata]++
	}

	senaraiKata := make([]string, 0, len(frekuensi))
	for kata := range frekuensi {
		senaraiKata = append(senaraiKata, kata)
	}
	sort.Strings(senaraiKata)
	fmt.Println("Ayat:", ayat)
	for _, kata := range senaraiKata {
		fmt.Printf("%s: %d kali\n", kata, frekuensi[kata])
	}
}
