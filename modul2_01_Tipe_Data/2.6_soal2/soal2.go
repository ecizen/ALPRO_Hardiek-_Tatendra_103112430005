package main

import "fmt"

func main() {

	 var (
		name string
		nim int
		prodi string
		kelas string
	 )

	 fmt.Print("Masukan nama anda :")
	 fmt.Scanln(&name)
	 fmt.Print("Masukan NIM anda :")
	 fmt.Scanln(&nim)
	 fmt.Print("masukan prodi anda :")
	 fmt.Scanln(&prodi)
	 fmt.Print("Masukan kelas anda :")
	 fmt.Scanln(&kelas)


	 fmt.Println("Perkenalkan saya adalah ", name , " salah satu mahasiswa Prodi", prodi  ,"dari kelas", kelas ,"dengan NIM", nim)


}