package main

import "fmt"

func main() {
	var tahun int

	fmt.Println("Masukan tahun nya")

	fmt.Scan(&tahun)

	if (tahun %4 == 0 && tahun%100 != 0) || tahun%400 == 0 {
		print("maka tahun ini adalah kabisat")
	} else {
		println("tahun ini bukan kabisat")
	}
}