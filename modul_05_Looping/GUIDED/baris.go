package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("masukan bilangan bulat 1: ")
	fmt.Scan(&a)
	fmt.Print("masukan bilangan bulat 2: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
}
