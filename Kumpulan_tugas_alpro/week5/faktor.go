package main

import "fmt"

func main() {
	var number int //dekelerasi number dengan tipe data int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number) // input number

	fmt.Printf("Faktor dari %d adalah:\n", number)
	for i := 1; i <= number; i++ { // perulangan index 1 dana number lebih kecil sama dengan index
		if number%i == 0 {
			fmt.Printf("%d true\n", i)  // mencetak hasil false or true
		} else {
			fmt.Printf("%d false\n", i) // Mencetak bukan faktor dan false
		}
	}
}
