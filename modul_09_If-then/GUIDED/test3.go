package main

import "fmt"

func main() {
	var input int

	
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&input)

	if input < 1000 || input > 9999 {
		fmt.Println("angka harus 4 digit")
		return
	}

	
	d1 := input / 1000            
	d2 := (input / 100) % 10      
	d3 := (input / 10) % 10       
	d4 := input % 10 
	
	
	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Println("terurut membesar")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Println("bilangan terurut mengecil")
	} else {
		fmt.Println("ini bilangan tidak terurut ")
	}
}
