package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		Ax, Ay, Bx, By, Cx, Cy float64
	)

	fmt.Print("masukan nila Ax dan Ay anda: ")
	fmt.Scan(&Ax, &Ay)

	fmt.Print("masukan nila Bx dan By anda: ")
	fmt.Scan(&Bx, &By)

	fmt.Print("masukan nila Cx dan Cy anda: ")
	fmt.Scan(&Cx, &Cy)

	// Menghitung panjang sisi segitiga
	AB := math.Sqrt(math.Pow(Bx-Ax, 2) + math.Pow(By-Ay, 2))
	BC := math.Sqrt(math.Pow(Cx-Bx, 2) + math.Pow(Cy-By, 2))
	CA := math.Sqrt(math.Pow(Ax-Cx, 2) + math.Pow(Ay-Cy, 2))

	// Menentukan sisi terpanjang
	maxSide := AB
	if BC > maxSide {
		maxSide = BC
	}
	if CA > maxSide {
		maxSide = CA
	}

	// Tampilkan hasil dengan dua angka di belakang koma
	fmt.Printf("%.2f\n", maxSide)
}
