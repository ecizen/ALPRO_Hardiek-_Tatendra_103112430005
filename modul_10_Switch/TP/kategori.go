package main

import (
	"fmt"
)

func main() {

	var usia int

	print("Masukan Usia Anda: ")
	fmt.Scan(&usia)

	switch {
		case usia >= 0 && usia <= 12:
		fmt.Println("Kategori: Anak-anak")
		case usia >= 13 && usia <= 17:
		fmt.Println("Kategori: Remaja")
		case usia >= 18 && usia <= 64:
		fmt.Println("Kategori: Dewasa")
		default:
			fmt.Println("Kategori: Lansia")
	}
}