package main

import (
	"fmt"

)

func main() {
	var (
		totalSemester, nilaiEprt int  //ini digunakan untuk mendeklarasikan total semseter dan niali eprt dengn int
		cekCumlaude             bool //mendekelerasikan cumlaude false or true
	)

	fmt.Println("Selamat datang di cek cumlaude by Hardiek")
	fmt.Print("Masukan jumlah semester yang anda lulus: ")
	fmt.Scan(&totalSemester) //input total Semester
	fmt.Print("Masukan nilai eprt anda: ")
	fmt.Scan(&nilaiEprt) // input nilai eprt

	cekCumlaude = totalSemester <= 8 && nilaiEprt >= 500 // cek cumlaude dengan boolean 

	println(cekCumlaude)

	if (cekCumlaude == true) { //penjelasan kenapa mahasiswa cumlaude apabila true
		println("Mahasiswa cumlaude dengan kuliah", totalSemester,"semester dan EPrt", nilaiEprt)
	} else if(totalSemester > 8){ //total semester diatas8 diakatakan tidak cumlaude maka akan print total semseter
	 	println("Mahasiswa tidak cumlaude karena kuliah hingga", totalSemester ,"smester")
	} else {
		println("Mahasiswa tidak cumlaude karena EPrt", nilaiEprt) //else digunakan selain diatas
	}

}