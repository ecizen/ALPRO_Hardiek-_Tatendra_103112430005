package main // main function

import (
	"fmt"
)

func main() {
	// menampilkan input dari variable
	var(
		a, b, c, d, e int
	)

	var hasil int

	fmt.Print("masukan input untuk angka ke a = ")
	fmt.Scanln(&a)
	fmt.Print("masukan input untuk angka ke b = ")
	fmt.Scanln(&b)
	fmt.Print("masukan input untuk angka ke c = ")
	fmt.Scanln(&c)
	fmt.Print("masukan input untuk angka ke d = ")
	fmt.Scanln(&d)
	fmt.Print("masukan input untuk angka ke e = ")
	fmt.Scanln(&e)

	hasil = a + b + c + d + e

	fmt.Println("hasilnya adalah =" , hasil)
	

}
