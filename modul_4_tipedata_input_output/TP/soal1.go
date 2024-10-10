package main


import (
    "fmt"
)

func main() {
	var pi , luas, keliling, jariJari float64

	pi = 3.14

	fmt.Print("masukan jari jari anda = ")
	fmt.Scan(&jariJari)

	luas = pi * jariJari * jariJari
	keliling =  2  * pi   * jariJari
	
	fmt.Printf("Luas lingkaran:      %.0f\n", luas)
	fmt.Printf("Keliling lingkaran:  %.0f\n", keliling)

}