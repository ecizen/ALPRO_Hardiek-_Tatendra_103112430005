package main

import "fmt"

func main() {
	var target, donasi, total, pendonasi int

	fmt.Scan(&target)

	for total < target {
		fmt.Scan(&donasi)

		total += donasi

		pendonasi++

		fmt.Printf("doantur %d: menyumbang %d. total terkumpul: %d\n", pendonasi, donasi, total)
	}
	fmt.Printf("target tercapai! total donasi: %d dari %d donatur.\n", total, pendonasi)
}