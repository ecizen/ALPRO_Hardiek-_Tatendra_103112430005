package main

import "fmt"

func main() {
	// dekalrasi a, b dan hasil dengan tipe data itneger
	var a, b, hasil int

	fmt.Print("masukan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("masukan nilai b: ")
	fmt.Scan(&b)

	// isi hasil dengan  nilai 1
	hasil = 1
	// memulai perulangan yang dimana batasnya adalah kurang dari masukan variable b
	for i := 0; i < b; i++ {
		// rumus perhitungan
		hasil *= a

	}
	// menampilkan hasil perpangkatan
	fmt.Print("hasil perpangkatan: ", hasil)

}
