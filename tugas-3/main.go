package main

import (
	"fmt"
	"strings"
)

func main() {
	// ==========================================
	// SOAL 1: Manipulasi String
	// ==========================================
	fmt.Println("--- SOAL 1 ---")
	kalimat := "halo halo bandung"
	angka := 2021

	// Mengganti kata "halo" menjadi "Hi"
	kalimatGanti := strings.ReplaceAll(kalimat, "halo", "Hi")

	// Menampilkan output sesuai format yang diminta
	fmt.Printf("\"%s\" - %d\n\n", kalimatGanti, angka)

	// ==========================================
	// SOAL 2: Kondisional (Indeks Nilai)
	// ==========================================
	fmt.Println("--- SOAL 2 ---")
	var nilaiJohn = 80
	var nilaiDoe = 50

	// Fungsi anonim untuk mengecek dan mencetak indeks agar kode lebih rapi
	cekIndeks := func(nama string, nilai int) {
		var indeks string
		if nilai >= 80 {
			indeks = "A"
		} else if nilai >= 70 && nilai < 80 {
			indeks = "B"
		} else if nilai >= 60 && nilai < 70 {
			indeks = "C"
		} else if nilai >= 50 && nilai < 60 {
			indeks = "D"
		} else {
			indeks = "E"
		}
		fmt.Printf("Nilai %s (%d) mendapatkan indeks: %s\n", nama, nilai, indeks)
	}

	cekIndeks("John", nilaiJohn)
	cekIndeks("Doe", nilaiDoe)
	fmt.Println()

	// ==========================================
	// SOAL 3: Switch Case (Format Tanggal)
	// ==========================================
	fmt.Println("--- SOAL 3 ---")
	var tanggal = 17
	var bulan = 8
	var tahun = 1945

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Printf("%d %s %d\n\n", tanggal, namaBulan, tahun)

	// ==========================================
	// SOAL 4: Kondisional (Generasi)
	// ==========================================
	fmt.Println("--- SOAL 4 ---")
	var tahunLahir = 1998 // Silakan ubah dengan tahun kelahiran Anda
	var generasi string

	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		generasi = "Baby boomer"
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		generasi = "Generasi X"
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Di luar rentang data yang diberikan"
	}

	fmt.Printf("Tahun kelahiran %d termasuk ke dalam: %s\n\n", tahunLahir, generasi)

	// ==========================================
	// SOAL 5: Looping dengan Syarat Tertentu
	// ==========================================
	fmt.Println("--- SOAL 5 ---")
	for i := 1; i <= 20; i++ {
		// Pengecekan harus dimulai dari kondisi yang paling spesifik terlebih dahulu
		// yaitu kelipatan 3 DAN ganjil
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 0 {
			// Kondisi genap
			fmt.Printf("%d - Berkualitas\n", i)
		} else {
			// Kondisi ganjil (sisanya)
			fmt.Printf("%d - Santai\n", i)
		}
	}
}
