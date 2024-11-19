package main

import "fmt"

func main() {
	var  usia int
	var kk bool

	fmt.Println("Selamat Datang.....>")
	fmt.Println("Masukan usia anda")
	fmt.Scan(&usia)

	fmt.Println("Masukan Ktp")
	fmt.Scan(&kk)

	if usia >= 17 && kk   {
		fmt.Println("Bisa membuat ktp")
	} else {
		fmt.Println("belum bisa membuat ktp")
	}
	fmt.Println("Program selesai")

}