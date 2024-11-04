package main

import "fmt"

func main() {
	var nilaiAnda int // deklerasikan variable nilai

	fmt.Scan(&nilaiAnda) // input nilai user

	if (nilaiAnda > 90) { // disini saya menggunakan else if untuk multiple condition
		println("Anda mendapatkan indeks A")
	} else if (nilaiAnda >= 80) { //setiap blok akan cek apakah nilai itu berpredikat apa jika true maka akan menjalankan
								  // block kode tersebut hingga menghasilkan nilai true atau sesuai 
		println("Anda mendapatkan indeks AB")
	} else if (nilaiAnda >= 70){
		println("Anda mendapatkan indeks B")
	} else {
		println("Anda mendapatkan indeks C")
	}
}