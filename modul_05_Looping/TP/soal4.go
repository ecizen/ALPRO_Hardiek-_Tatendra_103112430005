package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		bot := rand.Intn(100)
		var kesempatan = 5


		fmt.Println("Selamat datang di permainan tebak angka!")
		fmt.Println("Tebak angka antara 0 dan 99. Anda memiliki 3 kali percobaan.")

		for kesempatan > 0 {
			var player int

			fmt.Print("Masukkan jawaban anda: ")
			fmt.Scan(&player)

			if player == bot {
				fmt.Println("Jawaban anda benar! Selamat!")
				break
			} else if player < bot {
				fmt.Println("Jawaban anda terlalu rendah dari jawaban yang benar.")
			} else {
				fmt.Println("Jawaban anda terlalu tinggi dari jawaban yang benar.")
			}

			kesempatan--

			if kesempatan == 0 {
				fmt.Printf("Game selesai! Anda kalah. Jawaban yang benar adalah %d.\n", bot)
			}
		}

		var jawab string
		fmt.Print("Apakah Anda ingin bermain lagi? (ya/tidak): ")
		fmt.Scan(&jawab)

		if jawab != "ya" {
			fmt.Println("Terima kasih telah bermain!")
			break
		}
	}
}
