package main

import "fmt"

func main() {
	var (
		inputan, digit1, digit2, digit3  int
	)

	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&inputan)
	digit1 = inputan /100
	digit2 = (inputan % 100) / 10
	digit3 = inputan % 10


	fmt.Println(digit1 <= digit2 && digit2 <= digit3)

	// hasil = digit3*100 + digit2*10 + digit1

	// if (hasil < inputan) {
	// 	fmt.Println("false")
	// } else {
	// 	fmt.Println("true")
	// }

	

}