package main

import "fmt"

func main() {
	var (
		jam, menit, detik  int
	)

	fmt.Print("Masukan detik yang ingin dikonversi: ")
	fmt.Scan(&detik)

	jam =  detik / 3600
 	menit = (detik % 3600) / 60
 	detik = detik % 60

	fmt.Printf("%d jam, %d menit, %d detik\n", jam, menit, detik)

}