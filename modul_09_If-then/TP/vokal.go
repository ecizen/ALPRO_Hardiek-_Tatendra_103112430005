package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		masukanUser string // mendeklerasikan masukan user dengan tipe data string
	)

	fmt.Scan(&masukanUser) // digunakan untuk input user
	masukanUser = strings.ToUpper(masukanUser) // pada bagian ini string akan mengubah semua input menjadi kapital 


	// bagian dibawah ini cek apakah masukan user  sesua vokal jika true maka akan mengesekusi printah terdapat huruf vokal
	if masukanUser == "A" || masukanUser == "I" || masukanUser == "U" || masukanUser == "E" || masukanUser == "O" {
		fmt.Println("Terdapat huruf vokal")
	} else { // selain AUIEO akan mengesekusi printah ini
		fmt.Println("Huruf Konsonan")
	}
}