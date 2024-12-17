package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga string
	var pita string
	var jumlahBunga int

	for {

		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		fmt.Scanln(&bunga)

		if strings.ToLower(bunga) == "selesai" {
			break
		}

		if jumlahBunga > 0 {
			pita += " – "
		}

		pita += bunga
		jumlahBunga++
	}

	fmt.Printf("Pita: %s –\n", pita)
	fmt.Printf("Bunga: %d\n", jumlahBunga)
}
