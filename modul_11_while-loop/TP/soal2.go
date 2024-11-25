package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalBelanja float64
	var lanjutkan string


	for {
		var namaBarang string
		var harga float64

		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk menyelesaikan transaksi): ")
		fmt.Scanln(&namaBarang)

	
		if strings.ToLower(namaBarang) == "selesai" {
			break
		}

	
		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&harga)

		totalBelanja += harga
		fmt.Printf("Barang %s dengan harga %.2f telah ditambahkan.\n", namaBarang, harga)

		
		fmt.Print("Apakah Anda ingin menambahkan barang lagi? (ya/tidak): ")
		fmt.Scanln(&lanjutkan)

		
		if strings.ToLower(lanjutkan) == "tidak" {
			break
		}
	}

	fmt.Printf("\nTotal belanja Anda adalah: %.2f\n", totalBelanja)
	fmt.Println("Terima kasih telah berbelanja di toko kami!")
}
