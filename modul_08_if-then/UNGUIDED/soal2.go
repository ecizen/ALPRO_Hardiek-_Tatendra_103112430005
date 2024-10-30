package main

import "fmt"

func main() {

	var bilangan int
	var hasil string = "bukan" // Default hasil

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if (bilangan < 0 && bilangan%2 == 0) {
		hasil = "genap negatif"
	}


	fmt.Println(hasil)
}
