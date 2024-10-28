package main

import "fmt"

func main() {
	var (
		//mendeklarasikan variable luas, pi dan r dengan tipe data float32
		luas, keliling, pi, r float32
	)
	//Input jari jari user
	fmt.Print("masukan jari jari = ")
	fmt.Scan(&r)

	// memasukan niai phi 2 kome di belakang
	pi = 3.14
	// memasukan rumus luas dan keliling
	luas = pi * r * r
	keliling = 2 * pi * r
	// menampilkan hasil perhitungan luas dan keliling
	fmt.Println("luas lingkaran anda : ", luas)
	fmt.Println("keliling lingkaran anda : ", keliling)

}
