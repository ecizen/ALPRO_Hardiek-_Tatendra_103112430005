package main

import "fmt"

func main() {
	var x, y int
	var (
		hasil1, hasil2 bool
	)

	fmt.Scan(&x)
	fmt.Scan(&y)

	if y%x == 0 || x%y == 0 {
		hasil1 = (y % x == 0)
		hasil2 = (x % y == 0)
	}

	fmt.Println(hasil1)
	fmt.Println(hasil2)
}
