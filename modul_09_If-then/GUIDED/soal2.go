package main

import "fmt"

func main() {
	var  input string

	fmt.Print("Masukkan input: ")
	fmt.Scan(&input)


	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O"  {
		fmt.Println("huruf vokal")
	} else if (input >= "a" && input <= "z") || (input >= "A" && input <= "Z") {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}

	

}