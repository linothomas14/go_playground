package main

import "fmt"

func main() {
	fmt.Println(penjumlahan(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func penjumlahan(angka ...int) (total int) {

	for _, val := range angka {
		total += val
	}

	return
}
