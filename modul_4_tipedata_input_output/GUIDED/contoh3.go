package main

import "fmt"

func main() {

	var (
		berat, tinggi, bmi float32
	)

	fmt.Print("masukan berat badan anda: ")
	fmt.Scan(&berat)
	fmt.Print("masukan tinggi badan anda: ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Printf("hasil bmi anda adalah %.2F", bmi)

}