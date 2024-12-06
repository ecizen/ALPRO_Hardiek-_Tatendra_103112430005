package main

import "fmt"

func main() {
	var inputPassword ,Inputusername  string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&Inputusername)

	fmt.Print("Masukkan password: ")
	fmt.Scan(&inputPassword)

	var correctUsername string= "Admin"
	var correctPassword  string= "admin"

	jumlahPercobaanGagal := 0

	for inputPassword != correctPassword || Inputusername != correctUsername{
		fmt.Print("Masukkan username: ")
		fmt.Scan(&Inputusername)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&inputPassword)
		jumlahPercobaanGagal ++
		 
	}
	fmt.Println("Jumlah percobaan anda gagal", jumlahPercobaanGagal)

}