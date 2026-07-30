package main

import "fmt"

func main() {
	nama := "Kamat"
	umur := 99

	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)

	if umur >= 18 {
		fmt.Println("orangtua")
	} else {
		fmt.Println("bocil")
	}
}
