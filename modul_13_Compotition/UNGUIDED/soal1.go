package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan angka n: ")
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			count++
		}
	}

	fmt.Printf("Jumlah bilangan ganjil antara 1 hingga %d adalah %d\n", n, count)

}
