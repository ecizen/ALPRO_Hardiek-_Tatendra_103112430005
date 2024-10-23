package main

import "fmt"

func main() {
	
	var fals , qirat, dinar, dirham, n int

	//konversi qirat ke dinar dirham false
	fmt.Println("Masukan jumlah qirat nya")
	fmt.Scan(&qirat)

	dinar= qirat /600 % 10
	dirham = qirat  /60 % 10 
	fals = qirat /6 % 10
	qirat =  qirat % 6

	for i := 0; i < n; i++ {
		fmt.Print(dinar,dirham,fals, qirat)
	}

}