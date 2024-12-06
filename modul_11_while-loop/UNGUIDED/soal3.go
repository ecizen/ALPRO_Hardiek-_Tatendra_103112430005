package main

import "fmt"

func main() {

	var bilangan1, bilangan2 int

	fmt.Scan(&bilangan1)
	fmt.Scan(&bilangan2)

	if (bilangan2 == 0) {
		fmt.Println("Pembagi tidak boleh 0")
		return
	}

	hasilBagi := 0

	for bilangan1 >= bilangan2 {
		bilangan1 -= bilangan2
		hasilBagi ++
	}

	fmt.Println(hasilBagi)

}