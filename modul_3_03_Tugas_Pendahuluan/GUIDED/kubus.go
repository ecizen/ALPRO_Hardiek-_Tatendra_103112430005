package main

import "fmt"

func main() {
	var (
		volume, s int
	)

	fmt.Print("Masukan nilai sisinya = ")
	fmt.Scanln(&s)

	volume = s * s * s

	println("hasil keliling kubus adalah = ", volume, "m")

}

