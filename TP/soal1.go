package main

import "fmt"

func main() {
	
	var (
		n int
	)

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	if  n <= 0 {
		fmt.Println("Bilangan harus positif.")
        return
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*i) 
	}

}