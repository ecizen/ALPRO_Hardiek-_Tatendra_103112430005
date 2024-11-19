package main
import "fmt"

func main() {

// jawaban

//soal 1 ketika saya menginput 80.1 itu yg keluar adalah predikat D harusnya A
//soal 2 kesalahan program seharusnya menurt saya lebih efektif menggunakan multiple if menggunakan
//else if kenapa demikian karena program dapat membaca satu persatu nilai jika true maka akan di eksekusi
// jika false akan ke condisi berikutnya dan di program sebeelumnya terdapat nam="Predikat" ini type error karena nam float sedangkan valua
//nam sendiri string untuk print nilai  


// soal 3 ini hasil perbaikan sehingga ketika input dengan masukan: 93.5; 70.6; dan 49.5.keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’
 var nam float64
 var nmk string
 fmt.Print("Nilai akhir mata kuliah: ")
 fmt.Scan(&nam)

if nam > 80 {
	nmk="A"
} else if (nam > 72.5) {
	nmk="AB"
} else if (nam > 65) {
    nmk ="B"
} else if ( nam > 57.5){
	nmk="BC"
} else if (nam > 50) {
	nmk="C"
} else if (nam > 40) {
	nmk="D"
} else if (nam <= 40) {
	nmk="E"
}
fmt.Println("Nilai mata kuliah: ", nmk)
}
