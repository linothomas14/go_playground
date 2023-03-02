package cases

import "fmt"

func Cari_nilai_unik_dari_array_2D() {
	a := [][]int{ //Matrik 2D
		{0, 0, 1, 2},
		{2, 3, 3, 6, 1},
	}

	// posisi ke-3 dari 7

	uniqueVal := []int{}

	arr := parseTwoDimentionToOneDimentionArray(a)
	// lenArr := len(arr)
	for _, i := range arr {
		if !isDuplicate(arr, i) {
			uniqueVal = append(uniqueVal, i)
		}
	}

	fmt.Println(uniqueVal)
}

func parseTwoDimentionToOneDimentionArray(a [][]int) []int {

	arr := []int{}
	for _, i := range a {
		for _, j := range i {
			arr = append(arr, j)
		}
	}
	return arr
}

func isDuplicate(s []int, n int) bool {
	count := 0
	for _, v := range s {

		if v == n {
			count++
		}
	}

	if count > 1 {
		return true
	}

	return false
}
