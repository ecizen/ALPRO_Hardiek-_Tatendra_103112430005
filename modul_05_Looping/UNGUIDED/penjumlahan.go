package main

import "fmt"

func main() {
	// mendeklerasikan n dan hasil dengan tipe data int
	var n, hasil int

	// meminta inputan dari user
	fmt.Print("masukan nilai n anda: ")
	fmt.Scan(&n)

	// for looping 
	for i := 1; i <= n; i++ {
		hasil += i
	}
	// menampilkan hasil  
	fmt.Print("hasil penjumlahan: ", hasil)
}
