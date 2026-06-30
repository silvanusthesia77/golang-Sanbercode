package main

import "fmt"

// ==========================================
// SOAL 1
// ==========================================
type Mahasiswa struct {
	Nama string
	NIM  string
	Usia int
}

// ==========================================
// SOAL 2
// ==========================================
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// Method luas untuk segitiga (mengembalikan float64 karena alas*tinggi/2 bisa desimal)
func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas*s.tinggi)
}

// Method luas untuk persegi
func (p persegi) luas() int {
	return p.sisi * p.sisi
}

// Method luas untuk persegiPanjang
func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// ==========================================
// SOAL 3
// ==========================================
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Menggunakan pointer (*phone) agar perubahan pada slice 'colors' tersimpan ke objek asli
func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
}

// ==========================================
// SOAL 4
// ==========================================
type movie struct {
	title, genre   string
	duration, year int
}

var dataFilm = []movie{}

// Function menggunakan pointer ke slice (*[]movie) agar data di luar function ikut terupdate
func tambahDataFilm(title string, duration int, genre string, year int, b *[]movie) {
	baru := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}
	*b = append(*b, baru)
}

// ==========================================
// SOAL 5
// ==========================================
type Hewan interface {
	Suara() string
}

type Kucing struct{}

func (k Kucing) Suara() string {
	return "Meong"
}

type Anjing struct{}

func (a Anjing) Suara() string {
	return "Guk Guk"
}

// ==========================================
// MAIN FUNCTION
// ==========================================
func main() {
	// --- Eksekusi Soal 1 ---
	fmt.Println("=== JAWABAN SOAL 1 ===")
	mhs := Mahasiswa{
		Nama: "Budi Santoso",
		NIM:  "12345678",
		Usia: 20,
	}
	fmt.Printf("Nama: %s\nNIM: %s\nUsia: %d tahun\n\n", mhs.Nama, mhs.NIM, mhs.Usia)

	// --- Eksekusi Soal 2 ---
	fmt.Println("=== JAWABAN SOAL 2 ===")
	s := segitiga{alas: 10, tinggi: 5}
	p := persegi{sisi: 4}
	pp := persegiPanjang{panjang: 8, lebar: 5}

	fmt.Printf("Luas Segitiga: %.1f\n", s.luas())
	fmt.Printf("Luas Persegi: %d\n", p.luas())
	fmt.Printf("Luas Persegi Panjang: %d\n\n", pp.luas())

	// --- Eksekusi Soal 3 ---
	fmt.Println("=== JAWABAN SOAL 3 ===")
	myPhone := phone{
		name:   "iPhone 15",
		brand:  "Apple",
		year:   2023,
		colors: []string{"Black", "Blue"},
	}
	fmt.Println("Warna awal:", myPhone.colors)
	myPhone.tambahWarna("White") // Memanggil method pointer
	fmt.Println("Warna setelah ditambah:", myPhone.colors, "\n")

	// --- Eksekusi Soal 4 ---
	fmt.Println("=== JAWABAN SOAL 4 ===")
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// Perulangan untuk menampilkan output data film
	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n   duration: %d menit\n   genre: %s\n   year: %d\n", i+1, film.title, film.duration, film.genre, film.year)
	}
	fmt.Println()

	// --- Eksekusi Soal 5 ---
	fmt.Println("=== JAWABAN SOAL 5 ===")
	var hewan1 Hewan = Kucing{}
	var hewan2 Hewan = Anjing{}

	fmt.Println("Suara Kucing:", hewan1.Suara())
	fmt.Println("Suara Anjing:", hewan2.Suara())
}
