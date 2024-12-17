package main

import "fmt"

func main() {

	var bilangan int
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
