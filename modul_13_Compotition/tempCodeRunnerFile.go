package main

import "fmt"

// Fungsi untuk mengecek apakah suatu bilangan prima
func isBilanganPrima(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var limit int
    fmt.Print("Masukkan batas bilangan: ")
    fmt.Scanln(&limit)

    fmt.Println("Bilangan prima dari 1 hingga", limit, "adalah:")
    for i := 2; i <= limit; i++ {
        if isBilanganPrima(i) {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}
