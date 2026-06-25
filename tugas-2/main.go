package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	// tampilkan kalimat "Bootcamp Digital Skill Sanbercode Golang" yang tersusun dari gabungan variabel dalam setiap kata (5 variabel)

	bootcam := "Bootcamp"
	digital := "Digital"
	skil := " Skill"
	sanbercode := "Sanbercode"
	golang := " Golang"

	fmt.Println(bootcam, digital, skil, sanbercode, golang)

	// 	soal 2

	// terdapat variabel seperti di bawah ini:

	// halo := "Halo Dunia"

	// ganti kata "Dunia" menjadi "Golang" menggunakan packages strings
	halo := "Halo Dunia"
	find := "Dunia"
	replace := "Golang"
	newText := strings.Replace(halo, find, replace, 2)
	fmt.Println(newText)

	// 	soal 3

	// buatlah variabel-variabel seperti di bawah ini
	// gabungkan variabel-variabel tersebut agar menghasilkan output
	// saya Senang belajaR GOLANG

	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	kata2 := strings.ToUpper(string(kataKedua[0])) + kataKedua[1:]
	str := strings.ToUpper(kataKeempat)
	fmt.Println(kataPertama, kata2, kataKetiga, str)

	// 	soal 4

	// buatlah variabel-variabel seperti di bawah ini

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	intPertama, _ := strconv.Atoi(angkaPertama)
	iniKedua, _ := strconv.Atoi(angkaKedua)
	iniKetiga, _ := strconv.Atoi(angkaKetiga)
	iniKeempat, _ := strconv.Atoi(angkaKeempat)

	totalJumlah := intPertama + iniKedua + iniKetiga + iniKeempat
	fmt.Println(totalJumlah)

	// Soal 5

	// ubahlah variabel diatas menjadi integer dan jumlahkan semuanya
	// 1. Deklarasi variabel awal (string)
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// 2. Deklarasi variabel hasil (integer)
	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	// 3. Mengubah variabel string menjadi integer
	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	// 4. Melakukan operasi perhitungan matematika
	// Rumus Luas Persegi Panjang = panjang x lebar
	luasPersegiPanjang = panjang * lebar

	// Rumus Keliling Persegi Panjang = 2 x (panjang + lebar)
	kelilingPersegiPanjang = 2 * (panjang + lebar)

	// Rumus Luas Segitiga = (alas x tinggi) / 2
	luasSegitiga = (alas * tinggi) / 2

	// 5. Menampilkan hasil perhitungan
	fmt.Println("--- Hasil Perhitungan ---")
	fmt.Println("Luas Persegi Panjang     :", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang :", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga            :", luasSegitiga)
}
