package main


import (
    "fmt"
)

func main() {
	var pi , luas, keliling, jariJari float32

	pi = 3.14

	fmt.Print("masukan jari jari anda = ")
	fmt.Scan(&jariJari)

	luas = pi * jariJari * jariJari
	keliling =  2  * pi   * jariJari
	
	fmt.Print("luas lingkaran anda adalah = ")
	fmt.Print(luas, " dan " )

	fmt.Print("keliling lingkaran anda adalah = ")
	fmt.Printf("%.2f" , keliling)


}