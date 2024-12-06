package main

import "fmt"

func main() {
	var n, j int

	fmt.Println("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	j = n

	for j > 1 {
		fmt.Print(j, " X ")
		j = j -1
	}
	fmt.Println(1)
	

}