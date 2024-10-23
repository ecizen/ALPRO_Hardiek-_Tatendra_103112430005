package main

import "fmt"

func main() {
	var x, y, hasil int

	fmt.Println("Masukan nilai x: ")
	fmt.Scan(&x)
	fmt.Println("Masukan nilai y: ")
	fmt.Scan(&y)

	for i := x; i <= y; i++ {
		hasil +=i
	}
	fmt.Println(hasil) 
}