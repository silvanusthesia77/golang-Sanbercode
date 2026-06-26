package main

import "fmt"

func main() {
	// ==========================================
	// SOAL 1: Segitiga Pagar (#) dengan Looping
	// ==========================================
	fmt.Println("=== Soal 1 ===")
	tinggi := 7
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#") // Opsional, atau langsung print '#'
		}
		fmt.Println() // Pindah baris baru setelah satu baris selesai
	}

	// ==========================================
	// SOAL 2: Slice Slicing / Mengolah Array Kalimat
	// ==========================================
	fmt.Println("\n=== Soal 2 ===")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	// Mengambil elemen dari indeks ke-2 ("saya") sampai indeks terakhir
	var hasilSlice = kalimat[2:]
	fmt.Println(hasilSlice)

	// ==========================================
	// SOAL 3: Append Data ke Slice dan Looping dengan Angka
	// ==========================================
	fmt.Println("\n=== Soal 3 ===")
	var sayuran = []string{}

	// Menambahkan data menggunakan fungsi append
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	// Menampilkan dengan looping dan nomor urut (indeks + 1)
	for indeks, item := range sayuran {
		fmt.Printf("%d. %s\n", indeks+1, item)
	}

	// ==========================================
	// SOAL 4: Looping pada Map
	// ==========================================
	fmt.Println("\n=== Soal 4 ===")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	// Melakukan iterasi (looping) pada map untuk mengambil key dan value
	// Catatan: Urutan output pada map di Go bersifat acak (random) setiap kali dijalankan
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}

	// ==========================================
	// SOAL 5: Implementasi 3 Function Matematika/Geometri
	// ==========================================
	fmt.Println("\n=== Soal 5 ===")
	panjang := 12
	lebar := 4
	tinggiBalok := 8 // Menggunakan nama variabel tinggiBalok agar tidak bentrok dengan variabel 'tinggi' di Soal 1

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggiBalok)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}

// --- Defenisi Function untuk Soal 5 ---

// Fungsi menghitung luas persegi panjang: p * l
func luasPersegiPanjang(p, l int) int {
	return p * l
}

// Fungsi menghitung keliling persegi panjang: 2 * (p + l)
func kelilingPersegiPanjang(p, l int) int {
	return 2 * (p + l)
}

// Fungsi menghitung volume balok: p * l * t
func volumeBalok(p, l, t int) int {
	return p * l * t
}
