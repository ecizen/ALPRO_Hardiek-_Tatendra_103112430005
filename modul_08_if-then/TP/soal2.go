package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	if (bilangan % 2 == 0) {
		fmt.Println("Bilangan genap")
	} else {
		fmt.Println("Bilangan ganjil")
	}
}