package main

import (
	"fmt"
)

func main() {

	var (
		nilaiLulus    int
		nilaiAbsensi  int
		hitungLulus   bool
		hitungAbsensi bool
		hasilLulus    bool
	)

	nilaiLulus = 90
	nilaiAbsensi = 80

	hitungLulus = nilaiLulus > 80
	hitungAbsensi = nilaiAbsensi >= 80

	hasilLulus = hitungLulus && hitungAbsensi

	fmt.Println(hasilLulus)
}
