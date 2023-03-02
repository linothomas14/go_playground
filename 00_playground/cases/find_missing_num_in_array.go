package cases

import "fmt"

func Find_missing_num_in_array() {
	arr := []int{5, 4, 1, 3, 3, 3, 3, 3}
	result := []int{}
	max := findMax(arr)

	for i := 1; i < max; i++ {
		if !isExist(arr, i) {
			result = append(result, i)
		}
	}
	fmt.Println(result)

}

func findMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func isExist(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
