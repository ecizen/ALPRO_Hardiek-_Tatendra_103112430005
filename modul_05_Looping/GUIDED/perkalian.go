package main

import "fmt"

func main() {
	var a, b, hasil int

	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		hasil = hasil + a

	}
	fmt.Print("Hasil perhitungan: ", hasil)
}
