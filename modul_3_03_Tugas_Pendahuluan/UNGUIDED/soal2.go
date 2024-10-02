package main

import "fmt"

func main() {
	var (
		angkaPertama, angkaKedua float32 //mendeklerasikan inputan user angka 1 dan kedua
	)
	var operator string //dekelrasi operator 

	println("Selamat datang di kalkulator sederhana by Hardiek Tatendra")
	println("------------------------------------------------------------")

	print("Masukan angka pertama = ")
	fmt.Scan(&angkaPertama) //input angka pertama

	print("Masukan angka kedua = ")
	fmt.Scan(&angkaKedua) //input angka kedua

	println("Masukan angka dan operator yang valid =  +, -, *, /, %")
	fmt.Scan(&operator)  //input yang opreator

	switch operator { // swith case untuk memilih operasi
	case "+":
		fmt.Print("Hasil dari penjumlahan : ", angkaPertama, " + ", angkaKedua, "= ", angkaPertama+angkaKedua)
	case "-":
		fmt.Print("Hasil dari pengurangan : ", angkaPertama, " - ", angkaKedua, "= ", angkaPertama-angkaKedua)
	case "*":
		fmt.Print("Hasil dari perkalian : ", angkaPertama, " * ", angkaKedua, "= ", angkaPertama*angkaKedua)
	case "/":
		if angkaKedua  !=0 {
			fmt.Print("Hasil dari pembagiann : ", angkaPertama, " / ", angkaKedua, " = ", angkaPertama/angkaKedua)
		} else {
			fmt.Println("Operasi invalid karena angka kedua dalam pembagian tidak boleh dengan 0")
		}
	case "%":
		if angkaKedua != 0 {
			fmt.Print("Hasil dari penjumlahan modulus = " ,  int(angkaPertama) % int(angkaKedua))
		} else {
			fmt.Println("Operasi modulus tidak dapat dilakukan karena angka kedua adalah 0")
		}
	default:
		fmt.Println("Operator yang anda masukkan salah. Operator yang valid adalah +, -, *, /, %")
	}
}
