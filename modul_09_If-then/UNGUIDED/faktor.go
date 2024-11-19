package main

import (
	"fmt"
)

func main() {
	var b int


	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	
	isPrime := true
	for i := 2; i <= b/2; i++ {
		if b%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime && b > 1 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
