package main

import "fmt"


func cekBilanganSempurna(n int) string {

    sum := 0
    for i := 1; i < n; i++ {
        if n % i == 0 {
            sum += i
        }
    }

    if sum == n {
        return "Ya"
    }
    return "Tidak"
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    fmt.Println(cekBilanganSempurna(n))
}
