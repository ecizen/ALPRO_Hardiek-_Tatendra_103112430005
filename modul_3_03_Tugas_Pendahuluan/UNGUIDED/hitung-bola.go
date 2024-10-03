package main

import "fmt"

func main() {
	var (
		jariJari,
		pi,
		luas,
		volume float64

	)

	
	fmt.Println("Selamat datang di hitung luas dan volume bola")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Masukan jari-jari bola = ")
	fmt.Scan(&jariJari)
	
	pi = 3.1415926535
	luas = (pi * jariJari) * jariJari * 4

	volume = (4.0/3.0) *pi  * jariJari * jariJari * jariJari
	

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f", int(jariJari), volume, luas)



}