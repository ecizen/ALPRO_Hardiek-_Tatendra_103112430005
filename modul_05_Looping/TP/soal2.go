package main

import "fmt"

func main() {

	var input, i, j int

	fmt.Println("Masukan jumlah angka anda untuk membuat bintang")

	fmt.Scan(&input)
	
	for i = 0; i <= input; i++ {

		for j = 0; j <= i; j++ {
			print("*")
		}

		fmt.Println()
	}

}