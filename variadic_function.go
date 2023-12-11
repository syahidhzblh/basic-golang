package main

import "fmt"

func sumAll(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	fmt.Println(sumAll(10, 20, 30))
	fmt.Println(sumAll(20, 20, 60))

	//Jika sudah terlanjur ada slice, gunakan ... saat memanggil function
	umur := []int{21, 22, 23, 24, 25}
	fmt.Println(sumAll(umur...))
}
