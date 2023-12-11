package main

import (
	"fmt"
)

func main() {

	name := "Syahid"

	switch name {

	case "Syahid":
		fmt.Println("Halo Syahid")
	case "Hisbul":
		fmt.Println("Halo Hisbul")
	default:
		fmt.Println("Nama kamu siapa ?")
	}

	switch length := len(name); length > 5 {

	case true:
		fmt.Println("Nama Sudah Benar")
	case false:
		fmt.Println("Nama Belum Benar, kurang karakternya!")
	}

	name = ""
	length := len(name)

	switch {

	case length <= 10 && length > 1:
		fmt.Println("Nama Anda " + name)
		fmt.Println("Total Huruf Nama Anda", len(name), "Karakter")
		fmt.Println("Nama Sudah Benar")
	case length > 10:
		fmt.Println("Nama Anda " + name)
		fmt.Println("Total Huruf Nama Anda", len(name), "Karakter")
		fmt.Println("Nama Terlalu Panjang")
	default:
		fmt.Println("Nama Salah")
		fmt.Println("Total Huruf Nama Anda", len(name), "Karakter")
	}

}
