package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1
	// Closure = function tanpa nama

	var greeting = func(greet, name string) string {
		result := greet + ", " + name + "\n"
		return result
	}

	fmt.Printf("%s", greeting("Selamat pagi", "Yulyano Thomas Djaya"))

	// 2

	var max = 4
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	// 3
	// fungsi strings.Contains

	var res = strings.Contains("Yulyano", "ano")
	// true

	fmt.Printf("%v\n", res)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
