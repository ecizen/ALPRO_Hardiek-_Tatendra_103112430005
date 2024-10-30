package main

import "fmt"

func main() {
	var nilai int 

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if (nilai >= 70) {
		println("Nilai Ujian: ", nilai)
		println("Lulus")
	} else {
		println("Nilai Ujian: ", nilai)
        println("Tidak Lulus")
	}
}