package main

import "fmt"

func main() {

	var (
		jumlahJamKerja, upahPerJam float64
	) 


	fmt.Print("Masukan jumlah jam kerja dalam seminggu: ")
	fmt.Scan(&jumlahJamKerja)

	fmt.Print("Masukkan upah per jam: ")
    fmt.Scan(&upahPerJam)

	var jamNormal float64
	var jamLembur float64
	var totalGaji float64

	if (jumlahJamKerja > 40 ) {
		jamNormal = 40
		jamLembur = jumlahJamKerja - 40
	} else {
		jamNormal =  jumlahJamKerja 
		jamLembur = 0
	}

	totalGaji =  (jamNormal * upahPerJam + jamLembur * 1.5 * upahPerJam)

	var totalGajiBulanan float64 = 4 * totalGaji

	fmt.Printf("Total gaji anda seminggu ini adalah: %.2f", totalGaji)
	fmt.Print(" dan ")
	fmt.Printf("Total gaji anda dalam 1 bulan adalah: %.2f", totalGajiBulanan)



}