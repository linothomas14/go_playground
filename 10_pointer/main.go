package main

import (
	"fmt"
)

func main() {

	var a = 7
	var b = &a

	fmt.Printf("Nilai a = %d\n", a)
	fmt.Printf("Pointer a = %v\n", &a)
	fmt.Printf("Nilai b = %d\n", *b)
	fmt.Printf("Pointer b = %v\n", b)

	fmt.Println()

	*b = 10

	fmt.Printf("Nilai a = %d\n", a)
	fmt.Printf("Pointer a = %v\n", &a)
	fmt.Printf("Nilai b = %d\n", *b)
}
