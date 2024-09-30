package main

import (
	"fmt"
)

func main() {
	var (
		celcius, fahrenheit int
	)
	fmt.Print("Masukan suhu Fahrenheit anda: ")
	fmt.Scan(&fahrenheit)

	celcius =  (fahrenheit - 32) * 5/9

	
	fmt.Print("Hasil konversi suhu dari Fahrenheit ke Celsius adalah:  ", celcius)
}
