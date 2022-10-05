package main

import "fmt"

func main() {

	//1

	var positif uint8 = 255 //uint khusus angka positif saja
	fmt.Printf("Bilangan positif %d\n", positif)

	//2

	var desimal float32 = 3.14

	fmt.Printf("Angka desimal = %.3f\n", desimal)

	//3

	var isExist bool = true

	fmt.Printf("Apakah data A ada ? %t\n", isExist)

	var paragraf string
	paragraf = `Ini paragraf 1
ini paragraf 2
ini paragraf 3
nama saya "Yulyano Thomas Djaya"`

	fmt.Printf("%s\n", paragraf)
	fmt.Println(&paragraf, &isExist)
}
