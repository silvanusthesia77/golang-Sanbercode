package main

import (
	"fmt"
)

func main() {
	learnVariable()
	learnDataTypes()
}
func learnVariable() {
	// short declaratuion Variable
	name := "Silvanus Thesia"
	age := 28
	hobby := "football"
	fmt.Println("Name :", name)
	fmt.Println("Age :", age)
	fmt.Println("Hobby :", hobby)
	// Contoh Constant
	const appName = "Belajar Gratis"
	fmt.Println(appName)
}
func learnDataTypes() {
	var number int8 = 22
	angka := -2381918
	// Dibawah ini jadinya float64
	floatNilai := 1.0
	nilai := 0.123
	// dibawah ini adalah boolean
	isTrue := true
	isfalse := false
	// dibawah ini adalah string
	name := "Silvanus Thesia"
	fmt.Println(number)
	fmt.Println(angka)
	fmt.Println(nilai)
	fmt.Printf("Ini Float %v \n", floatNilai)
	fmt.Println(isTrue == isfalse)
	fmt.Println(isTrue != isfalse)
	fmt.Println(name)
}
