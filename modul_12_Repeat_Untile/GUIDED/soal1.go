package main

import "fmt"

func main() {
	var kata string
	var jumlah int

	fmt.Scan(&jumlah)
	fmt.Scan(&kata)


	counter :=0 

	for done := false; !done;  {
		fmt.Println(kata)
		counter++
 
		done = (counter >= jumlah)
		
	}
}