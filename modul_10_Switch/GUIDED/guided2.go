package main

import "fmt"

func main() {
	var nama string

	fmt.Scan(&nama)

	switch nama {
		case "nepenthes":
			fmt.Println("Termasuk Tanaman Karnivora")
			fmt.Println("Asli indonesia")
		case "venus":
			fmt.Println("Termasuk Tanaman Karnivora")
			fmt.Println("Bukan asli Indonesia")
		case "tillandsia":
			fmt.Println("Termasuk Tanaman Karnivora")
			fmt.Println("Bukan asli Indonesia")
		default:
			fmt.Println("Input tidak valid")
	}
	

}