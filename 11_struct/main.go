package main

import "fmt"

type Student struct {
	nama   string
	alamat string
}

type Keluarga struct {
	kepala_keluarga string
	anak            []Anak
}

type Anak struct {
	nama string
}

func main() {

	s0 := Student{}
	s0.nama = "Budi"

	s1 := Student{
		nama:   "Yulyano Thomas Djaya",
		alamat: "Bogor",
	}

	s2 := Student{"Wahyu", "Jakarta"}

	fmt.Printf("Nama s0 : %v\n", s0.nama)
	fmt.Printf("Nama s1 : %v\n", s1.nama)
	fmt.Printf("Nama s2 : %v\n", s2.nama)
	fmt.Println()

	// 2

	anak_sejuk := []Anak{
		{nama: "Mumpuni Nur Izzati"},
		{nama: "Yulyano Thomas Djaya"},
		{nama: "Arumi Kharisma Cendani"},
	}
	sejuk_family := Keluarga{
		kepala_keluarga: "Sejuk Karyanto",
		anak:            anak_sejuk,
	}
	fmt.Println(sejuk_family)
	fmt.Printf("Kepala keluarga dari keluarga sejuk_family : %s\n", sejuk_family.kepala_keluarga)
	fmt.Printf("Memiliki anak berjumlah : %d\n", len(sejuk_family.anak))
	fmt.Printf("Terdiri dari : \n")

	for _, anak := range anak_sejuk {
		fmt.Println(anak)
	}
	fmt.Println()

	// 3 Anonymous Struct

	avanza := struct {
		nama      string
		warna     string
		kecepatan string
	}{}
	avanza.nama = "Avanza"
	avanza.warna = "Silver"
	avanza.kecepatan = "140 KM/jam"

	fmt.Println(avanza)
}
