package main

import (
	"fmt"
)

func main() {
	var beratKantongKiri, beratKantongKanan float64

	for {

		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKantongKiri, &beratKantongKanan)

		if beratKantongKiri < 0 || beratKantongKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat := beratKantongKiri + beratKantongKanan
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisihBerat := beratKantongKiri - beratKantongKanan
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}

		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
