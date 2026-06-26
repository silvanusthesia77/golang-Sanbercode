package main

import (
	"fmt"
	"math"
	"strings"
)

// ==========================================
// SOAL 1: Function dengan Named Return Value
// ==========================================
// Menggunakan 'kalimat' sebagai named return value
func introduce(nama, jk, profesi, umur string) (kalimat string) {
	panggilan := "Pak"
	if jk == "perempuan" {
		panggilan = "Bu"
	}

	// Format string sesuai dengan output yang diharapkan
	kalimat = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggilan, nama, profesi, umur)

	// Karena menggunakan named return value, cukup panggil return
	return
}

// ==========================================
// SOAL 2: Variadic Function
// ==========================================
func buahFavorit(nama string, buah ...string) string {
	var buahQuotes []string

	// Menambahkan tanda kutip pada setiap item buah
	for _, b := range buah {
		buahQuotes = append(buahQuotes, fmt.Sprintf(`"%s"`, b))
	}

	// Menggabungkan slice menjadi satu string dipisahkan koma
	listBuah := strings.Join(buahQuotes, ", ")
	namaLower := strings.ToLower(nama)

	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", namaLower, listBuah)
}

// ==========================================
// SOAL 4: Function Factorial
// ==========================================
// Menggunakan tipe data uint64 agar kapasitas penyimpanan angka lebih besar
func factorial(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	var hasil uint64 = 1
	var i uint64
	for i = 2; i <= n; i++ {
		hasil *= i
	}

	return hasil
}

// ==========================================
// SOAL 5: Pointer pada Parameter Function
// ==========================================
// Parameter luas dan keliling menggunakan pointer (*) untuk mengubah nilai aslinya
func hitungLingkaran(luas *float64, keliling *float64, r float64) {
	*luas = math.Pi * math.Pow(r, 2) // Rumus Luas = π * r^2
	*keliling = 2 * math.Pi * r      // Rumus Keliling = 2 * π * r
}

func main() {
	// --- Eksekusi Soal 1 ---
	fmt.Println("=== Soal 1 ===")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// --- Eksekusi Soal 2 ---
	fmt.Println("\n=== Soal 2 ===")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	// --- Eksekusi Soal 3 ---
	fmt.Println("\n=== Soal 3 ===")
	var dataFilm = []map[string]string{}

	// Pembuatan Closure Function
	tambahDataFilm := func(judul, durasi, genre, tahun string) {
		film := map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}

	// Memasukkan data melalui closure
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	// --- Eksekusi Soal 4 ---
	fmt.Println("\n=== Soal 4 ===")
	fmt.Println("factorial(5) menghasilkan", factorial(5))
	fmt.Println("factorial(7) menghasilkan", factorial(7))

	// --- Eksekusi Soal 5 ---
	fmt.Println("\n=== Soal 5 ===")
	var luasLingkaran float64
	var kelilingLingkaran float64
	var jariJari float64 = 7.0

	// Mempassing reference (alamat memori) menggunakan operator '&'
	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, jariJari)

	fmt.Printf("Luas Lingkaran     : %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran : %.2f\n", kelilingLingkaran)
}
