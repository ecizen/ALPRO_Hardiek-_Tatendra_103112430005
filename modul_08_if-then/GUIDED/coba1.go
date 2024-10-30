package main

import "fmt"

func main() {
	var bilangan int

	fmt.Println("masukan bilangan bulat")
	fmt.Scan(&bilangan)

	if (bilangan < 0) {
		bilangan = -bilangan
		fmt.Println(bilangan)
	} else {
		fmt.Println(bilangan)
	}
}