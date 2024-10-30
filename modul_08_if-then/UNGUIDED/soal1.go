package main

import "fmt"

func main() {
	var jumlahOrang int // mendeklerasikan jumlah orang


	fmt.Print("Masukan jumlah orang: ")
	fmt.Scan(&jumlahOrang)

	var jumlahMotor int

	if (jumlahOrang%2 == 0) {
		 jumlahMotor = jumlahOrang / 2
	}

	if (jumlahOrang%2 != 0) {
		 jumlahMotor = (jumlahOrang / 2) + 1
	}

	fmt.Println("jadi total motor yang diperlukan: ", jumlahMotor)

}