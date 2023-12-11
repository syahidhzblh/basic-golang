package main

import "fmt"

func forLoop() {

	//For Cara Ribet

	z := 1

	for z <= 10 {

		fmt.Println("Perulangan ke ", z)
		z++
	}

	fmt.Println("Perulangan Pertama Selesai")

	//For Common Usual
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("Perulangan ke-2 Selesai")

	//For Range
	names := []string{"Syahid", "Hizbullah", "Ganteng"}

	for index, isi_index := range names {
		fmt.Println("index", index, "=", isi_index)
	}

	//For Range (Hanya Ambil Valuenya saja, tanpa index)
	for _, isi_index := range names {
		fmt.Println(isi_index)
	}
}
