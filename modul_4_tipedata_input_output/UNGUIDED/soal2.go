package main

import "fmt"

func main() {
	var (
		tokg, tinggi, bmi float32
	)

	fmt.Print("Masukan Bmi anda: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukan tinggi anda: ")
	fmt.Scan(&tinggi)

	tokg =  bmi * (tinggi * tinggi)

	fmt.Printf("jadi berat badan anda adalah: %.0f", tokg)
}