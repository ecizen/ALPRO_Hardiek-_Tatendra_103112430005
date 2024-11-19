package main

import "fmt"

func main() {

	var gram, Totalbiaya, hargaKg, biayaTambahan int

	fmt.Scan(&gram)

	kg := gram / 1000
	gr := gram % 1000

	perKg := 10000

	hargaKg = perKg * kg
	

	fmt.Println("Berat parsel (gram): ", gram)
	fmt.Println("Detail berat: ",kg,"kg", "+", gr, "gr")

	if (gr >= 500) {
		
		biayaTambahan = 5 * gr
		Totalbiaya =  hargaKg + biayaTambahan
		
		fmt.Println("Detail biaya: Rp.", hargaKg, "+", biayaTambahan) 
		fmt.Println("Total biaya: Rp.",Totalbiaya)
		
	} else if  (gr < 500) {
		
		biayaTambahan = 15 * gr
		Totalbiaya =  hargaKg + biayaTambahan

		fmt.Println("Detail biaya: Rp.", hargaKg, "+", biayaTambahan) 
		fmt.Println("Total biaya: Rp.",Totalbiaya)
	} else {
		Totalbiaya =  hargaKg + biayaTambahan

		fmt.Println("Detail biaya: Rp.", hargaKg, "+", biayaTambahan) 
		fmt.Println("Total biaya: Rp.",Totalbiaya)
		
	}

	fmt.Println("End Program")

}