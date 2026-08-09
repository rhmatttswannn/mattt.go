package main

import "fmt"

func main() {
	// ============================================
	// TOPIK 2: OPERATOR & IF / ELSE
	// ============================================

	fmt.Println("========== OPERATOR MATEMATIK ==========")
	a := 10
	b := 3

	fmt.Println("a =", a, ", b =", b)
	fmt.Println("Tambah  a + b =", a+b)
	fmt.Println("Tolak   a - b =", a-b)
	fmt.Println("Kali    a * b =", a*b)
	fmt.Println("Bahagi  a / b =", a/b)   // hasil integer
	fmt.Println("Modulo  a % b =", a%b)   // baki pembahagian

	// Operator perbandingan
	fmt.Println("\n========== OPERATOR PERBANDINGAN ==========")
	fmt.Println("a == b :", a == b) // sama?
	fmt.Println("a != b :", a != b) // tidak sama?
	fmt.Println("a > b  :", a > b)  // lebih besar?
	fmt.Println("a < b  :", a < b)  // lebih kecil?
	fmt.Println("a >= b :", a >= b)
	fmt.Println("a <= b :", a <= b)

	// ============================================
	// IF / ELSE
	// ============================================
	fmt.Println("\n========== IF / ELSE ==========")

	umur := 17

	// Asas if/else
	if umur >= 18 {
		fmt.Println("Anda DEWASA, boleh memandu!")
	} else {
		fmt.Println("Anda BAWAH UMUR, belum boleh memandu.")
	}

	// if / else if / else
	markah := 75

	fmt.Printf("\nMarkah: %d → ", markah)
	if markah >= 90 {
		fmt.Println("Gred A - Cemerlang!")
	} else if markah >= 80 {
		fmt.Println("Gred B - Baik!")
	} else if markah >= 70 {
		fmt.Println("Gred C - Memuaskan")
	} else if markah >= 60 {
		fmt.Println("Gred D - Perlu usaha lagi")
	} else {
		fmt.Println("Gred E - Gagal")
	}

	// Operator logik: && (DAN), || (ATAU), ! (TIDAK)
	fmt.Println("\n========== OPERATOR LOGIK ==========")
	ada_ic := true
	umur2 := 20

	if ada_ic && umur2 >= 18 {
		fmt.Println("Boleh mengundi! (ada IC DAN cukup umur)")
	}

	hujan := false
	ada_payung := true

	if hujan || ada_payung {
		fmt.Println("Boleh keluar (hujan ATAU ada payung)")
	}

	// Nested if (if dalam if)
	fmt.Println("\n========== NESTED IF ==========")
	jantina := "lelaki"
	umur3 := 25

	if jantina == "lelaki" {
		if umur3 >= 18 {
			fmt.Println("Lelaki dewasa")
		} else {
			fmt.Println("Kanak-kanak lelaki")
		}
	} else {
		if umur3 >= 18 {
			fmt.Println("Perempuan dewasa")
		} else {
			fmt.Println("Kanak-kanak perempuan")
		}
	}

	// Switch - alternatif kepada if/else berganda
	fmt.Println("\n========== SWITCH ==========")
	hari := "Isnin"

	switch hari {
	case "Sabtu", "Ahad":
		fmt.Println(hari, "→ Hujung minggu! Rehat 🎉")
	case "Isnin":
		fmt.Println(hari, "→ Awal minggu, semangat! 💪")
	case "Jumaat":
		fmt.Println(hari, "→ Hampir hujung minggu! 😊")
	default:
		fmt.Println(hari, "→ Hari biasa, teruskan kerja!")
	}
}
