package main

import "fmt"

func main() {
	 var ( farenheitValue , to_kelvin float32)

	 
	 fmt.Println("Selamat data di konversi suhu dari Farenheit ke Kelvin")
	 fmt.Println("------------------------------------------------------")
	 fmt.Print("Masukan suhu dalam satuan Farenheit = ")
	 
	 fmt.Scan(&farenheitValue)
	 to_kelvin = (farenheitValue - 32) * 0.5/0.9 + 273.15

	 fmt.Printf("Selamat, hasil konversi %.2f F ke %.2f K\n", farenheitValue, to_kelvin)



}