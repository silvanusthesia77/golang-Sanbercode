package main

import (
	"fmt"
	"math"
	"strings"
)

// ==========================================
// SOAL 1: STRUCT & INTERFACE BANGUNAN
// ==========================================

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Method Soal 1
func (s segitigaSamaSisi) luas() int     { return (s.alas * s.tinggi) / 2 }
func (s segitigaSamaSisi) keliling() int { return 3 * s.alas }

func (p persegiPanjang) luas() int     { return p.panjang * p.lebar }
func (p persegiPanjang) keliling() int { return 2 * (p.panjang + p.lebar) }

func (t tabung) volume() float64        { return math.Pi * math.Pow(t.jariJari, 2) * t.heightSim() }
func (t tabung) heightSim() float64     { return t.tinggi } // helper nama variabel pembantu jika konflik
func (t tabung) luasPermukaan() float64 { return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi) }

func (b balok) volume() float64 { return float64(b.panjang * b.lebar * b.tinggi) }
func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

// ==========================================
// SOAL 2: STRUCT & INTERFACE PHONE
// ==========================================

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type displayer interface {
	showData() string
}

func (p phone) showData() string {
	colorList := strings.Join(p.colors, ", ")
	return fmt.Sprintf("Nama: %s\nBrand: %s\nTahun: %d\nWarna: %s", p.name, p.brand, p.year, colorList)
}

// ==========================================
// SOAL 3: EMPTY INTERFACE FUNCTION
// ==========================================

func luasPersegi(sisi int, status bool) interface{} {
	if sisi == 0 {
		if status {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if status {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

// ==========================================
// SOAL 5: SIMULASI PACKAGE MATEMATIKA
// ==========================================
// Karena digabung satu file, ini merepresentasikan fungsi dari package matematika

func MatematikaTambah(a int, b int) int {
	return a + b
}

func MatematikaKali(a int, b int) int {
	return a * b
}

// ==========================================
// MAIN FUNCTION (EKSEKUSI SEMUA SOAL)
// ==========================================

func main() {
	// --- EKSEKUSI SOAL 1 ---
	fmt.Println("========= JAWABAN SOAL 1 =========")
	s := segitigaSamaSisi{alas: 6, tinggi: 4}
	p := persegiPanjang{panjang: 5, lebar: 3}
	t := tabung{jariJari: 7, tinggi: 10}
	b := balok{panjang: 4, lebar: 3, tinggi: 2}

	var bd hitungBangunDatar
	bd = s
	fmt.Printf("Segitiga -> Luas: %d, Keliling: %d\n", bd.luas(), bd.keliling())
	bd = p
	fmt.Printf("Persegi Panjang -> Luas: %d, Keliling: %d\n", bd.luas(), bd.keliling())

	var br hitungBangunRuang
	br = t
	fmt.Printf("Tabung -> Volume: %.2f, Luas Permukaan: %.2f\n", br.volume(), br.luasPermukaan())
	br = b
	fmt.Printf("Balok -> Volume: %.2f, Luas Permukaan: %.2f\n\n", br.volume(), br.luasPermukaan())

	// --- EKSEKUSI SOAL 2 ---
	fmt.Println("========= JAWABAN SOAL 2 =========")
	myPhone := phone{
		name:   "iPhone 13 Pro",
		brand:  "Apple",
		year:   2021,
		colors: []string{"Sierra Blue", "Graphite", "Gold", "Silver"},
	}
	var d displayer = myPhone
	fmt.Println(d.showData())
	fmt.Println()

	// --- EKSEKUSI SOAL 3 ---
	fmt.Println("========= JAWABAN SOAL 3 =========")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
	fmt.Println()

	// --- EKSEKUSI SOAL 4 ---
	fmt.Println("========= JAWABAN SOAL 4 =========")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Type Assertion
	textPrefix := prefix.(string)
	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)

	total := angka1[0] + angka1[1] + angka2[0] + angka2[1]
	output := fmt.Sprintf("%s%d + %d + %d + %d = %d", textPrefix, angka1[0], angka1[1], angka2[0], angka2[1], total)
	fmt.Println(output)
	fmt.Println()

	// --- EKSEKUSI SOAL 5 ---
	fmt.Println("========= JAWABAN SOAL 5 =========")
	// Menggunakan simulasi fungsi matematika
	hasilTambah := MatematikaTambah(10, 5)
	hasilKali := MatematikaKali(10, 5)

	fmt.Printf("Hasil Tambah (10 + 5): %d\n", hasilTambah)
	fmt.Printf("Hasil Kali (10 * 5): %d\n", hasilKali)
	fmt.Println("==================================")
}
