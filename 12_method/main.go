package main

import (
	"fmt"
)

type Burung struct {
	nama  string
	suara string
}

type Person struct {
	nama string
}

func main() {

	// 1

	kenari := Burung{
		nama:  "Kenari",
		suara: "Cuit.. cuit..",
	}
	elang := Burung{
		nama:  "Elang",
		suara: "Arkk.. Arkk..",
	}
	kenari.berkicau()
	elang.berkicau()
	fmt.Println()

	// 2

	p1 := Person{
		nama: "Yulyano Thomas Djaya",
	}

	fmt.Println("Nilai awal", p1.nama)
	p1.ganti_nama("Lino Thomas")
	fmt.Println("Nilai kedua", p1.nama)
	p1.ganti_nama2("Ano")
	fmt.Println("Nilai ketiga", p1.nama)
}

func (b Burung) berkicau() {
	fmt.Printf("Burung %s berkicau %s\n", b.nama, b.suara)
}

func (p Person) ganti_nama(nama string) {
	fmt.Println("Berubah menjadi", nama)
	p.nama = nama
}

func (p *Person) ganti_nama2(nama string) {
	fmt.Println("Berubah menjadi", nama)
	p.nama = nama
}
