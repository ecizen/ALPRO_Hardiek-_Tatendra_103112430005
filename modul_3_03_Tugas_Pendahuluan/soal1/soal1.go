package main

import "fmt"

func main() {
	 var (
		s int
		hitungLuas int
		hitungKeliling int //mendeklariskan sebuah variabele s, hitung luas dan keliling
	 )

	 s = 27 //sisi sudah ditentukan dari study case

	 hitungLuas = s * s //memberi nilai /rumushitung hitung luas 
	 hitungKeliling = 4 * s //memberi nilai /rumushitung hitung keliling

	 fmt.Println("Hasil hitung luas alun alun purwokerto adalah = ", hitungLuas, "m") //print hitung luas
	 fmt.Println("Hasil hitung keliling alun alun purwokerto adalah = ", hitungKeliling) //print hitung keliling


}