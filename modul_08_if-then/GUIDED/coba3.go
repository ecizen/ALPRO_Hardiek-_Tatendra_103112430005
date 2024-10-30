package main

import "fmt"

func main() {
	var bilangan1 int
	var bilangan2 bool
	fmt.Scan(&bilangan1)

	bilangan2 = false

	if (bilangan1 > 0 && bilangan1%2 ==0) {
		bilangan2 = true
	} 

	fmt.Println(bilangan2)
}