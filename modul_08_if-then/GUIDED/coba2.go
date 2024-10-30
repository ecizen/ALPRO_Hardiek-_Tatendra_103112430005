package main

import "fmt"

func main() {
	var a
	var b string

	fmt.Scan(&a)

	b = "positif"

	if (b > 0) {
		fmt.Println("positif")
	}

	fmt.Println(b)
}