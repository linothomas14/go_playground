package main

import "fmt"

func main() {

	// 1 - slice

	var animals = []string{"cat", "dog", "horse", "bird", "fish", "goat", "cow", "bug"}
	fmt.Println(animals)
	fmt.Println(animals[2:5]) // [horse bird fish]
	fmt.Println(animals[:3])  // [cat dog horse]

	fmt.Println(&animals[0]) // pointer dari setiap item
	fmt.Println(&animals[1])
	fmt.Println(&animals[5])

	// 2 - akses item menggunakan 2 index

	a_animals := animals[0:3]
	a_animals[1] = "anjing" // mengganti item akan mengganti slice asal karena reference yang sama

	fmt.Println(animals[1])

	// 3 - append

	a_animals = append(a_animals, "burung", "ikan") // append akan berdampak kepada slice asal juga

	// fmt.Println(animals[5])
	fmt.Println(a_animals)
	fmt.Println(cap(a_animals))
	fmt.Println(animals)

	fmt.Println("cap animals sebelum: ", cap(animals))
	fmt.Println("len animals sebelum: ", len(animals))

	animals = append(animals, "dino")

	fmt.Println("cap animals sesudah: ", cap(animals))
	fmt.Println("len animals sesudah: ", len(animals))

	fmt.Println(&animals[7])
	fmt.Println(&animals[8])
	fmt.Println(animals)

	// 4 - copy

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2
}
