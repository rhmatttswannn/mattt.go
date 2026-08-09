package main

import "fmt"

func main() {
	// ============================================
	// TOPIK 3: LOOP (PERULANGAN)
	// ============================================
	// Dalam Go, hanya ada satu keyword loop: 'for'
	// Tapi boleh digunakan dalam banyak cara!

	// --- CARA 1: for biasa (macam for dalam bahasa lain) ---
	fmt.Println("========== FOR BIASA ==========")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Bilangan ke-%d\n", i)
	}

	// --- CARA 2: for seperti while ---
	fmt.Println("\n========== FOR SEPERTI WHILE ==========")
	kiraan := 1
	for kiraan <= 5 {
		fmt.Println("Kiraan:", kiraan)
		kiraan++
	}

	// --- CARA 3: for selamanya (infinite loop) ---
	// Guna 'break' untuk keluar
	fmt.Println("\n========== FOR INFINITE + BREAK ==========")
	nombor := 0
	for {
		nombor++
		fmt.Println("Nombor:", nombor)
		if nombor == 3 {
			fmt.Println("Berhenti di nombor 3!")
			break // keluar dari loop
		}
	}

	// --- CONTINUE: skip ke iterasi seterusnya ---
	fmt.Println("\n========== CONTINUE (skip nombor genap) ==========")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip nombor genap
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- FOR BERSARANG (Nested Loop) ---
	fmt.Println("\n========== NESTED LOOP (Jadual Darab) ==========")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// --- Contoh praktikal: Bintang piramid ---
	fmt.Println("\n========== PIRAMID BINTANG ==========")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// --- Contoh: FizzBuzz (soalan interview klasik!) ---
	fmt.Println("\n========== FIZZBUZZ (1 - 20) ==========")
	// Jika boleh dibahagi 15 → FizzBuzz
	// Jika boleh dibahagi 3  → Fizz
	// Jika boleh dibahagi 5  → Buzz
	// Selainnya → nombor biasa
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
