package main

import (
	"fmt"
)

func main() {
	// Variable dan Constant
	text := "hello world"
	text = "Apa kabar dunia"
	fmt.Println(text)

	// Constant
	const appGuard = "Applikasi Garuda"
	fmt.Println(appGuard)

	// Tipe Data
	var number int8 = 123
	negativeNumber := -123
	decimal := 2.312
	fmt.Println(number)
	fmt.Println(negativeNumber)
	fmt.Println(decimal)
	isRaining := true
	fmt.Println(isRaining)
	decnum := 2.44
	fmt.Printf("Bilangan : %.3f\n", decnum)
	fmt.Printf("Bilangan : %f\n", decnum)
	message := "Halo Bhy"
	fmt.Printf("Message : %v", message)
	name := "wanus thesia"
	age := 22
	hobby := "Sport"
	fmt.Printf("Nama saya %s , usi %d hobby saya %v\n", name, age, hobby)

	// Konversi Data Menggunakan Teknik Casting

	var a float64 = 44
	fmt.Println(a)
}
