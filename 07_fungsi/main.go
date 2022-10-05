package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// 1 Fungsi menggabungkan string

	namas := []string{"Yulyano", "Thomas", "Djaya"}
	value := gabung("Haloo", namas)
	fmt.Println(value)

	// 2 Fungsi angka random

	rand.Seed(time.Now().Unix())
	for i := 0; i <= 5; i++ {
		fmt.Println(getRandNum(0, 100))
	}

	// 3 Fungsi luas dan keliling Lingkaran

	jari_jari := 14.0
	luas, keliling := hitung_lingkaran(jari_jari)
	fmt.Printf("Jari jari Lingkaran : %.f\n", jari_jari)
	fmt.Printf("Luas : %.2f\n", luas)
	fmt.Printf("Keliling : %.2f\n", keliling)

}

func gabung(msg string, arr []string) string {

	result := msg + " " + strings.Join(arr, " ")

	return result
}

func getRandNum(min, max int) int {
	var value int

	value = rand.Intn(max-min) + min

	return value
}

func hitung_lingkaran(jari_jari float64) (luas float64, kel float64) {

	var PHI float64
	if int(jari_jari)%7 == 0 {
		PHI = 22 / 7
	} else {
		PHI = math.Pi
	}

	luas = PHI * math.Pow(jari_jari, 2)
	kel = 2 * PHI * jari_jari

	return
}
