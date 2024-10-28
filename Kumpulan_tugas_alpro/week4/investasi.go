package main

import "fmt"

type investasiFormula struct {
	totalInvestasiAwal float64
	totalPenjualan     float64
	keuntunganKotor    float64
	biayaTransaksi     float64
	pajakKeuntungan    float64
	keuntunganBersih   float64 //pengugunaan struck untuk formula
}

const persenBiayaTransaksi = 0.002 //bagian iniadalah constant yang dimana nilai tidak bisa di ubah 0.2%
const persenPajakKeuntungan = 0.1    //bagian iniadalah constant yang dimana nilai tidak bisa di ubah 10%

func main() {
	var rumus investasiFormula
	var (
		hargaBeli, hargaJual, jumlahSaham float64 //deklerasikan variable dengan tipedata float64
	)

	// Input dari user
	fmt.Print("Masukkan Harga Beli: Rp ")
	fmt.Scan(&hargaBeli)
	fmt.Print("Masukkan Harga Jual: Rp ")
	fmt.Scan(&hargaJual)
	fmt.Print("Masukkan Jumlah Saham: ")
	fmt.Scan(&jumlahSaham)

	rumus.totalInvestasiAwal = hargaBeli * jumlahSaham // rumus investasi awak
	rumus.totalPenjualan = hargaJual * jumlahSaham //rumus total penjualan
	rumus.keuntunganKotor = rumus.totalPenjualan - rumus.totalInvestasiAwal //rumus keuntungan kotor
	rumus.biayaTransaksi = rumus.totalPenjualan * persenBiayaTransaksi // rumus biaya transaksi
	rumus.pajakKeuntungan = rumus.keuntunganKotor * persenPajakKeuntungan // rumus pajak 
	if rumus.pajakKeuntungan < 0 { //menambahkan logika jika pajak min atau dibawah 0 maka pajak keuntunganya Rp 0
		rumus.pajakKeuntungan = 0 
	}
	rumus.keuntunganBersih = rumus.keuntunganKotor - rumus.biayaTransaksi - rumus.pajakKeuntungan // keuntunganBersih

	// Menampilkan hasil dari user
	fmt.Println("\nInformasi Investasi Saham:")
	fmt.Printf("Harga Beli: Rp %.2f\n", hargaBeli)
	fmt.Printf("Harga Jual: Rp %.2f\n", hargaJual)
	fmt.Printf("Jumlah Saham: %.0f\n", jumlahSaham)
	fmt.Printf("Total Investasi Awal: Rp %.2f\n", rumus.totalInvestasiAwal)
	fmt.Printf("Total Penjualan: Rp %.2f\n", rumus.totalPenjualan)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", rumus.keuntunganKotor)
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", rumus.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan: Rp %.2f\n", rumus.pajakKeuntungan)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", rumus.keuntunganBersih)
}
