package main

import "fmt"

func main() {

	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	fmt.Println(bulan)

	slice1 := bulan[:3]
	slice2 := bulan[0:6]
	slice3 := bulan[6:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
