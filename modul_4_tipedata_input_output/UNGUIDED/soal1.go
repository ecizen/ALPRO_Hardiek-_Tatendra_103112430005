package main

import "fmt"

func main() {
	var (
		totalBelanja, harga, discount int
	)

	fmt.Scan(&harga)
	fmt.Scan(&discount)

	totalBelanja = harga - (harga * discount / 100)

	fmt.Println(totalBelanja)



}