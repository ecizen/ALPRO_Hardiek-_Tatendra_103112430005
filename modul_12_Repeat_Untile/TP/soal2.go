package main

import (
	"fmt"
	
)

func main() {

	var userInput string

	var jawaban string


	jawaban = "telkom"


	for userInput != jawaban {

		fmt.Print("Masukkan jawaban anda: ")
		fmt.Scan(&userInput)

		if userInput == jawaban {
			fmt.Println("Hello", userInput)
			break
		}

	}

}