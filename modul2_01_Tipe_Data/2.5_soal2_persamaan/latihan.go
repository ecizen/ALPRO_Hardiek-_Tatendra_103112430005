package main

import "fmt"

func main() {
	var (
		X float32
	)

	var hasil float32

	fmt.Print("masukan nilai ke x =")
	fmt.Scanln(&X)

	hasil = (2/(X + 2)) + 5

	
	fmt.Println(hasil)

}
