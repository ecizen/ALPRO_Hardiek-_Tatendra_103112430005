package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var player int

	var jawaban int

	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(10) + 1

	jawaban = randomNumber


	fmt.Println("Selamat datang di permainan tebak angka!")

	for player != jawaban {

		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&player)

		if player == jawaban {
			fmt.Println("Jawaban anda benar! Selamat!")
			break
		}  else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}

	}

}