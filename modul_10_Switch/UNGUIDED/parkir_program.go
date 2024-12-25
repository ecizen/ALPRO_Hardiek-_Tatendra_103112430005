package main

import "fmt"

func main() {
	var kendaraan string
	var durasi float64
	var tarif int

	fmt.Print("Masukkan jenis kendaraan (motor, mobil, truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	waktu := int(durasi)
	switch {
	case durasi > float64(waktu):
		waktu++
	case waktu < 1:
		waktu = 1
	}

	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	biaya := waktu * tarif
	fmt.Printf("Jenis kendaraan: %s\n", kendaraan)
	fmt.Printf("Durasi Waktu: %d jam\n", waktu)
	fmt.Printf("Biaya Parkir: Rp%d\n", biaya)
}