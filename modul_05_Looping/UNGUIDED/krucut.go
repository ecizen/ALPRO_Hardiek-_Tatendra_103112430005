package main

import "fmt"

func main() {
	// deklarasikan variable n dengan tipe data integer
	var n int

	fmt.Print("masukan nilai n anda")
	fmt.Scan(&n) //meminta inputan user

	// looping untuk menghitung volume krucut
	for i := 1; i <= n; i++ {
		var r, h float64
		fmt.Scan(&r)
		fmt.Scan(&h)
		vol := (1.0 / 3.0) * 3.14 * r * r * h
		// memunculkan hasil 
		fmt.Printf("Luas segitiga ke-%d = %.2f\n", i, vol)
	}

}
