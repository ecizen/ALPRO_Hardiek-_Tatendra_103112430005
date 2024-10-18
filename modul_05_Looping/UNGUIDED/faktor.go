package main

import "fmt"

func main() {
	// deklarasikan variable  n dengan tipe data int
	var (
		n , hasil int
	) 
	
	fmt.Print("masukan nilai n: ")
	
	// meminta input user
	fmt.Scan(&n)
	// menggunakan logika if jika n kurang dari 0 adalah bilangan negativ
	if n < 0 {
		// bila bukan positif maka akan muncul peringatan dibawah
		fmt.Println("harus positif")
	}

	// menetapkan nilai hasil = 1
	hasil = 1
	// memulai looping untuk faktorial
	for i := 0; i < n; i++ {
		// rumus perhitungan
		hasil *= i + 1
	}
	// menampilkan hasil 
	fmt.Printf("faktorial %d adalah: %d\n", n, hasil)

}
