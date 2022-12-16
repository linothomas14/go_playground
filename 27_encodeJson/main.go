package main

import (
	"encoding/json"
	"fmt"
)

type Mahasiswa struct {
	NPM     string `json:"npm"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

func main() {
	var Data []Mahasiswa

	jsonString := `[
		{"nama": "Yulyano Thomas Djaya", "npm": "56419764", "jurusan": "Teknik Informatika"},
		{"nama": "Waldo Felix", "npm": "1222122", "jurusan": "Teknik Informatika"},
		{"nama": "Muhammad Rifqi Alfurqon", "npm": "4141414", "jurusan": "Teknik Informatika"}
		]
		`

	jsonData := []byte(jsonString)

	err := json.Unmarshal(jsonData, &Data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, val := range Data {
		fmt.Printf("Mahasiswa ke-%d\n", i+1)
		fmt.Printf("NPM:%s\n", val.NPM)
		fmt.Printf("Nama:%s\n", val.Nama)
		fmt.Printf("Jurusan:%s\n\n", val.Jurusan)
	}

}
