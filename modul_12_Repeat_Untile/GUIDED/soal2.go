package main

import (
	"fmt"

)

func main() {
	var bilangan int


	for done := true; done; {

		fmt.Print("masukan bilangan :")
		fmt.Scan(&bilangan)
		
		done = (bilangan <= 0)
	}
	fmt.Println(bilangan, "adalah bilangan positif")
}