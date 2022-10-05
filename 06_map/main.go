package main

import "fmt"

func main() {
	weeks := map[int]string{
		1: "senin",
		2: "selasa"}

	weeks[3] = "rabu"
	weeks[5] = "jumat"
	weeks[4] = "kamis"
	weeks[6] = "sabtu"
	weeks[7] = "minggu"

	fmt.Println(weeks)

	for key, value := range weeks {
		fmt.Printf("%s hari \tke-%d dalam seminnggu\n", value, key)
	}
	fmt.Println()

	// 2 delete function

	nama := "Yulyano Thomas Djaya"
	barisan := map[int]string{
		1: "halo",
		2: "nama saya",
		3: nama,
	}

	fmt.Println("sebelum :", barisan)
	delete(barisan, 2)
	fmt.Println("sesudah :", barisan)
	fmt.Println()

	// 3 isExist

	mahasiswa := map[int]string{
		1: "Yulyano",
		2: "Thomas",
		3: "Djaya",
	}

	value, isExist := mahasiswa[1]

	if isExist {
		fmt.Printf("Mahasiswa %s, hadir\n", value)
	} else {
		fmt.Printf("Mahasiswa yang disebutkan,tidak hadir\n")
	}
	fmt.Println()
	// 4 slice and map

	data := []map[string]string{
		{"nama": "Yulyano", "kelas": "4IA01", "hobi": "gaming"},
		{"nama": "Thomas", "kelas": "4IA01"},
		{"nama": "Djaya", "kelas": "4IA01", "hobi": "reading"},
	}

	for _, value := range data {
		fmt.Println(value["hobi"])
	}

}
