package cases

import "fmt"

func Pisahkan_tahun_kabisat() {
	arr := []int{40, 50, 104, 200, 300, 400, 800, 1496, 1997, 2004, 2005}

	result := []int{}

	for _, i := range arr {
		// 100,400

		if i%100 == 0 {
			if i%400 == 0 {
				result = append(result, i)
			}
			continue
		}

		if i%4 == 0 {
			result = append(result, i)
		}

	}

	fmt.Println(result)
}
