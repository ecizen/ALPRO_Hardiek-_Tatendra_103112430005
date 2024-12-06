package main

import "fmt"

func main() {

	var token string


	fmt.Print("Masukkan token: ")
	fmt.Scan(&token)


	for token != "user1234" {
		fmt.Print("Masukkan token: ")
		fmt.Scan(&token)
	}

	fmt.Println("Selamat anda berhasil login")

	
	
}