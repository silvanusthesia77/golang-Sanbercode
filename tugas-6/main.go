package main

import "fmt"

// Fungsi introduce menggunakan pointer untuk mengubah isi variabel sentence
func introduce(sentence *string, name string, gender string, profession string, age string) {
	title := "Pak"
	if gender == "perempuan" {
		title = "Bu"
	}
	
	// Mengubah nilai yang ditunjuk oleh pointer sentence
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, profession, age)
}

func main() {
	var sentence string
	
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}