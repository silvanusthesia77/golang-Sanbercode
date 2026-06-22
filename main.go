package main

import (
	"fmt"
)

func main() {
	latihan()
	name := "Silvanus Thesia"
	age := 28
	hobby := "FootBall"
	fmt.Printf("Nama Saya %s ,  saya berusia %d tahun , hobby saya adalah %s \n", name, age,hobby )
}
func latihan (){
	var name = "John Doe"
var age = "28"
sentence := `Nama Saya Dalah "` + name + `usia saya "` + age + `"`
fmt.Println()
fmt.Println(sentence)


}