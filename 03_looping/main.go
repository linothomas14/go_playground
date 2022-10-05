package main

import "fmt"

func main() {

	// 1

	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			fmt.Printf("%d ganjil\n", i)
		} else {
			fmt.Printf("%d genap\n", i)
		}
	}
	fmt.Println()

	// 2
	// 3 Dag , 5 Dig , 5 dan 3 Der

	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d der\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d dag\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d dig\n", i)
		} else {
			fmt.Println(i)
		}
	}

	// 3

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j <= i {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}

	// 4

	for i := 10; i > 0; i-- {
		for j := 9; j > 0; j-- {
			if j >= i {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}

	fmt.Println("")

	// 5

	for i := 10; i > 0; i-- {
		temp := 1
		for j := 9; j > 0; j-- {
			if temp < i {
				fmt.Print(j)
			}

			temp++
		}
		fmt.Println()
	}

	// 6

	for i := 10; i > 0; i-- {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

}
