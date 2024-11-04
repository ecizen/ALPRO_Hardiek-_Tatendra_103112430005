package main

import "fmt"

func main() {
	
	var umur int // mendeklerasikan umur int
	var kewarganegaraan string // mendeklerasikan kewarganegaraan string

	fmt.Scan(&umur) // inputan user umur dan kewarganegaraan
	fmt.Scan(&kewarganegaraan)

	if (umur >= 17 && kewarganegaraan == "Indonesia"){ // cek apakah kedua syarat terpenuhi jika iya anda bisa mengikuti pemilu
		fmt.Println("Anda bisa mengikuti Pemilu")
	} else if (umur < 17 && kewarganegaraan != "Indonesia") { // cek apakah kedua syarat terpenuhi jika tidak akan memberi penjelasan umur anda belum cukup dan anda bukan warga Indoensia
		fmt.Println("Anda kurang cukup umur dan bukan asli Indonesia untuk mengikuti Pemilu")
	} else if (umur < 17 ) { // cek apakah umur kurang dari 17 jika true maka akan mengekseskusi Anda belum cukup umur
		fmt.Println("Anda belum cukup umur")
	} else { // cek selain diatas akan menampilkan anda bukan
		fmt.Println("Anda bukan asli Indonesia")
	}
}