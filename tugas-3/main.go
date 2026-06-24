package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

// Struct untuk Soal 2
type Device struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

// Struct untuk Soal 3
type Book struct {
	Title       string
	Desc        string
	Author      string
	ReleaseYear int
}

// Function untuk Soal 4
func hitungTabung(jariJari float64, tinggi float64) (volume float64, luasAlas float64, kelilingAlas float64) {
	pi := math.Pi

	luasAlas = pi * math.Pow(jariJari, 2)
	kelilingAlas = 2 * pi * jariJari
	volume = luasAlas * tinggi

	return volume, luasAlas, kelilingAlas
}

func main() {
	// Data JSON untuk Soal 1 dan 2
	var jsonString = `[
		{"vendor":"Acer","model":"beTouch E110"},
		{"vendor":"Samsung","model":"Galaxy S III"},
		{"vendor":"Samsung","model":"Galaxy Note II"},
		{"vendor":"Sony","model":"Xperia Z"},
		{"vendor":"Sony","model":"Ericsson J100"}
	]`

	fmt.Println("========================================")

	// soal 1
	var dataMap []map[string]string
	err := json.Unmarshal([]byte(jsonString), &dataMap)
	if err != nil {
		log.Fatalf("Error decode JSON (Soal 1): %v", err)
	}

	fmt.Println("--- Hasil Soal 1 (Vendor Samsung) ---")
	for _, data := range dataMap {
		if data["vendor"] == "Samsung" {
			fmt.Printf("Vendor: %s, Model: %s\n", data["vendor"], data["model"])
		}
	}

	fmt.Println("\n========================================")

	// soal 2
	var dataStruct []Device
	err = json.Unmarshal([]byte(jsonString), &dataStruct)
	if err != nil {
		log.Fatalf("Error decode JSON (Soal 2): %v", err)
	}

	fmt.Println("--- Hasil Soal 2 (Vendor Sony) ---")
	for _, data := range dataStruct {
		if data["vendor"] == "Sony" {
			fmt.Printf("Vendor: %s, Model: %s\n", data["vendor"], data["model"])
		}
	}

	fmt.Println("\n========================================")

	// soal 3
	books := []Book{
		{Title: "Belajar Golang", Desc: "Buku dasar pemrograman Go", Author: "Budi", ReleaseYear: 2021},
		{Title: "Clean Code", Desc: "Panduan menulis kode yang rapi", Author: "Robert C. Martin", ReleaseYear: 2008},
		{Title: "The Pragmatic Programmer", Desc: "Tips menjadi programmer handal", Author: "Andrew Hunt", ReleaseYear: 1999},
		{Title: "Sistem Informasi Manajemen", Desc: "Konsep dasar SIM", Author: "Andi", ReleaseYear: 2019},
		{Title: "Desain Database", Desc: "Cara merancang database relasional", Author: "Siti", ReleaseYear: 2022},
	}

	booksJSON, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Fatalf("Error encode JSON (Soal 3): %v", err)
	}

	fmt.Println("--- Hasil Soal 3 (Encode ke JSON) ---")
	fmt.Println(string(booksJSON))

	fmt.Println("\n========================================")

	// soal 4
	// Setup route untuk web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jariJari := 7.0
		tinggi := 10.0

		volume, luasAlas, kelilingAlas := hitungTabung(jariJari, tinggi)

		result := fmt.Sprintf("jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f",
			jariJari, tinggi, volume, luasAlas, kelilingAlas)

		fmt.Fprint(w, result)
	})

	fmt.Println("--- Hasil Soal 4 (Web Server) ---")
	fmt.Println("Server berjalan di port 8080. Buka http://localhost:8080 di browser Anda.")

	// Menjalankan web server (akan mem-block eksekusi kode di bawahnya)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server gagal dijalankan: %v", err)
	}
}
