package main

import "fmt"

func main() {
	var warna1, warna2, warna3 string

	warnaBenar1 := "merah"
	warnaBenar2 := "kuning"
	warnaBenar3 := "hijau"

	var percobaan int = 0
	var berhasil bool = true

	for percobaan < 5 {
		fmt.Println("Percobaan ke", percobaan+1)
		fmt.Print("Masukkan warna 1: ")
		fmt.Scanln(&warna1)
		fmt.Print("Masukkan warna 2: ")
		fmt.Scanln(&warna2)
		fmt.Print("Masukkan warna 3: ")
		fmt.Scanln(&warna3)

		if warna1 != warnaBenar1 || warna2 != warnaBenar2 || warna3 != warnaBenar3 {
			berhasil = false
		}

		percobaan++
	}

	fmt.Printf("Berhasil: %v\n", berhasil)
}
