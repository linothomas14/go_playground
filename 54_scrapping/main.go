package main

import (
	"fmt"
	"scrapping/scrapping"
	"time"
)

func main() {
	start := time.Now()
	scrapping.ScrapBAAK()
	elapsed := time.Since(start) // menghitung waktu eksekusi
	fmt.Printf("Waktu eksekusi: %s\n", elapsed)
}

// document.querySelectorAll(table#example) <---- to get all items ..
