package main

import "fmt"

func main() {

	// 1
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var nama string = "Yulyano Thomas Djaya"
	var alamat string
	alamat = "Sentul, Bogor"

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	universitas := "Universitas Gunadarma"

	fmt.Printf("Halo !! nama saya %s,\nsaya tinggal di %s,\nsaya mahasiswa %s\n\n", nama, alamat, universitas)

	// 2

	var buah, hewan, lokasi string
	buah = "mangga"
	hewan = "kucing"
	lokasi = "Jakarta"

	fmt.Printf("Buah = %s\nHewan = %s,Lokasi = %s\n", buah, hewan, lokasi)

	// 3 Variable underscore

	/*
		Tidak bisa digunakan/ditampilkan,
		berguna untuk menampung nilai balik yang tidak digunakan)
	*/

	name, _ := "john", "wick"
	fmt.Println(name)

	// 4 Variable pointer

	var point = new(string)
	*point = "Yulyano Thomas Djaya"
	fmt.Println(point)
	fmt.Println(*point)
}
