package main

import "fmt"

func main() {

	var (
		jmlahBarang, totalPoint, bonus int
	)

	fmt.Scan(&jmlahBarang)

	totalPoint = 0
	bonus = 5

	for i := 1; i <= jmlahBarang; i++ {
		if i <=5 {
			totalPoint += 10
		} else {
			totalPoint += 10 +  bonus
		}
	}
	fmt.Printf("Poin yang didapatkan: %d poin\n", totalPoint)

}