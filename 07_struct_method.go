package main

import "fmt"

type Pelajar struct {
	Nama   string
	Umur   int
	Markah int
}

func (pelajar Pelajar) Perkenalan() string {
	return fmt.Sprintf("Saya %s, umur %d tahun", pelajar.Nama, pelajar.Umur)
}

func (pelajar Pelajar) Lulus() bool {
	return pelajar.Markah >= 60
}

func (pelajar Pelajar) UbahMarkahSalinan(markahBaru int) {
	pelajar.Markah = markahBaru
	fmt.Println("Markah di dalam value receiver:", pelajar.Markah)
}

func (pelajar *Pelajar) UbahMarkah(markahBaru int) bool {
	if markahBaru < 0 || markahBaru > 100 {
		return false
	}
	pelajar.Markah = markahBaru
	return true
}

func main() {
	fmt.Println("========== TOPIK 7: STRUCT DAN METHOD ==========")
	fmt.Println("Struct menggabungkan beberapa field dengan jenis yang boleh berbeza.")
	fmt.Println("Method ialah fungsi dengan receiver, iaitu nilai yang menerima panggilan.")

	fmt.Println("\n========== MEMBUAT STRUCT ==========")
	ali := Pelajar{
		Nama:   "Ali",
		Umur:   18,
		Markah: 80,
	}
	fmt.Printf("Data Ali: %+v\n", ali)
	fmt.Println("Nama:", ali.Nama)
	fmt.Println("Markah:", ali.Markah)

	fmt.Println("\n========== MENGUBAH FIELD ==========")
	ali.Umur = 19
	fmt.Println("Umur selepas diubah:", ali.Umur)

	fmt.Println("\n========== ZERO VALUE ==========")
	var pelajarBaru Pelajar
	fmt.Printf("Nama: %q | Umur: %d | Markah: %d\n",
		pelajarBaru.Nama, pelajarBaru.Umur, pelajarBaru.Markah)
	fmt.Println("Field string bermula dengan string kosong, field int bermula dengan 0.")

	fmt.Println("\n========== MEMANGGIL METHOD ==========")
	fmt.Println(ali.Perkenalan())
	fmt.Println("Ali lulus:", ali.Lulus())

	fmt.Println("\n========== VALUE RECEIVER ==========")
	ali.UbahMarkahSalinan(40)
	fmt.Println("Markah Ali selepas method dipanggil:", ali.Markah)
	fmt.Println("Receiver Pelajar menerima salinan; perubahan field int tidak mengubah asal.")

	fmt.Println("\n========== POINTER RECEIVER ==========")
	berjaya := ali.UbahMarkah(90)
	fmt.Println("Perubahan diterima:", berjaya)
	fmt.Println("Markah Ali sekarang:", ali.Markah)
	fmt.Println("Receiver *Pelajar merujuk data asal, jadi perubahan kekal.")
	fmt.Println("Go mengambil alamat ali secara automatik ketika memanggil method ini.")

	berjaya = ali.UbahMarkah(120)
	fmt.Println("Markah 120 diterima:", berjaya)
	fmt.Println("Markah Ali kekal:", ali.Markah)

	fmt.Println("\n========== MENYALIN STRUCT ==========")
	salinan := ali
	salinan.Nama = "Budi"
	fmt.Println("Nama asal:", ali.Nama)
	fmt.Println("Nama salinan:", salinan.Nama)
	fmt.Println("Field contoh ini disalin sebagai nilai, berbeza dengan map sebelumnya.")
	fmt.Println("Jika field berisi map atau slice, salinan masih boleh berkongsi data di dalamnya.")

	fmt.Println("\n========== PRAKTIKAL: SENARAI PELAJAR ==========")
	senarai := []Pelajar{
		ali,
		{Nama: "Siti", Umur: 18, Markah: 95},
		{Nama: "Budi", Umur: 19, Markah: 45},
	}
	fmt.Println("Tambah 10 markah untuk pelajar yang belum lulus.")
	for indeks := range senarai {
		if !senarai[indeks].Lulus() {
			senarai[indeks].UbahMarkah(senarai[indeks].Markah + 10)
		}
	}
	fmt.Println("Gunakan indeks untuk mengubah elemen slice, bukan salinan dari range.")

	jumlahLulus := 0
	for _, pelajar := range senarai {
		status := "Belum lulus"
		if pelajar.Lulus() {
			status = "Lulus"
			jumlahLulus++
		}
		fmt.Printf("%s | Umur: %d | Markah: %d | %s\n",
			pelajar.Nama, pelajar.Umur, pelajar.Markah, status)
	}
	fmt.Printf("Jumlah lulus: %d daripada %d pelajar\n", jumlahLulus, len(senarai))
}
