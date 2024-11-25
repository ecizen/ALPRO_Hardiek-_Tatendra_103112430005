package main

import "fmt"

func main() {

	var password = "user1234"

	var maxPercobaan = 3
	var percobaan = 0 


	for percobaan < maxPercobaan {
		var inputPassword string
		fmt.Print("Masukkan password: ")
		fmt.Scan(&inputPassword)

		
		if inputPassword == password {
			fmt.Println("Login berhasil")
			return 
		} else {
			percobaan++ 
			fmt.Printf("Password salah. Kesempatan tersisa: %d\n", maxPercobaan-percobaan)
		}
	}


	fmt.Println("Login ditolak")
}
