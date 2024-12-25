package main

import "fmt"

func main() {
	var input int
	var hasil int
	var berikutnya int
	fmt.Scan(&input)

	switch {
	case input%10 == 0:
		hasil = input / 10
		fmt.Printf("Kategori: Bilangan kelipatan 10\nHasil pembagian antara: %d / 10 = %d\n", input, hasil)
	case input%5 == 0 && input != 5:
		hasil = input * input
		fmt.Printf("Kategori: Bilangan kelipatan 5\nHasil perpangkatan: %d * %d = %d\n", input, input, hasil)
	case input%2 == 0:
		berikutnya = input + 1
		hasil = input * berikutnya
		fmt.Printf("Kategori: Bilangan genap\nHasil perkalian: %d * %d = %d\n", input, berikutnya, hasil)
	default:
		berikutnya = input + 1
		hasil = input + berikutnya
		fmt.Printf("Kategori: Bilangan ganjil\nHasil perkalian: %d + %d = %d\n", input, berikutnya, hasil)
	}
}