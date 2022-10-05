package main

import "fmt"

func main() {

	// 1

	var arr = [...]string{
		"a",
		"b",
		"c",
		"d",
	}

	fmt.Println(arr)
	fmt.Printf("Jumlah item array arr adalah %d\n", len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Printf("index ke %d : %s\n", i, arr[i])
	}

	for i, ar := range arr {
		fmt.Printf("index ke %d : %s\n", i, ar)

	}

	// 2

	var number1 = [2][3]int{{3, 2, 3}, {1, 2, 3}}
	var number2 = [3][5]int{{9, 8, 7, 6}, {1, 2, 3}}

	fmt.Println(number1)
	fmt.Println(number2)
}
