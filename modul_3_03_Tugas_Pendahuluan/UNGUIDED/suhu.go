package main

import "fmt"

func main() {
	var (
		celcius, farenheit, reamur, kelvin float64
	)

	fmt.Println("Masukan celcius nya")
	fmt.Scan(&celcius)

	farenheit = (celcius * 9 / 5) + 32
	reamur = celcius * 4 / 5
	kelvin = (farenheit + 459.67) * 5 / 9

	fmt.Printf("Temperatur celcius: %.0f\n", celcius)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", farenheit)
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)

}
