package main

import (
	"fmt"
)

func main() {
	var harga int

	fmt.Println("Masukan Harga Awal barang:")
	fmt.Scan(&harga)
	
	for harga > 0 {
		fmt.Println("Masukkan harga tambahan atau 0 untuk berhenti:")
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		harga += input
		
		fmt.Println("Updated harga:", harga)
	}

	fmt.Println("Total Harga Akhir:", harga)
}
