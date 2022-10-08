package library

import "fmt"

func Sapa(nama string) {
	fmt.Println("Halo")
	perkenalkan(nama)
}

func perkenalkan(nama string) {
	fmt.Println("Perkenalkan nama saya " + nama)
}

func init() {
	fmt.Println("--> library/library.go imported")
}
