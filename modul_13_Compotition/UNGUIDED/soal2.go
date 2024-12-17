package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan angka n: ")
	fmt.Scan(&n)

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(n, "bukan bilangan prima")
			return
		}
	}
	fmt.Println(n, "adalah bilangan prima")
}
