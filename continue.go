package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("Perulangan Selesai")

	angka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	total_angka_ganjil := 0
	total_angka_genap := 0
	total_angka_prima := 0

	//Cari Angka Ganjil
	fmt.Println("Bilangan Ganjilnya Adalah")
	for _, isi_angka := range angka {

		if isi_angka%2 == 1 {
			fmt.Println(isi_angka)
			total_angka_ganjil++
		}

	}
	fmt.Println("Total Bilangan Ganjilnya adalah", total_angka_ganjil)

	//Cari Angka Genap
	fmt.Println("Bilangan Ganjilnya Adalah")
	for _, isi_angka := range angka {

		if isi_angka%2 == 0 {
			fmt.Println(isi_angka)
			total_angka_genap++
		}

	}
	fmt.Println("Total Bilangan Genapnya adalah", total_angka_genap)

	//Cari Angka Prima
	fmt.Println("Bilangan Primanya Adalah")
	for _, isi_angka := range angka {

		if    {
			
		}else{
			fmt.Println(isi_angka)
			total_angka_prima++
		}

	}
	fmt.Println("Total Bilangan primanya adalah", total_angka_prima)
}
