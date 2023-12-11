package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runningApplication() {

	//Defer function akan dipanggil setelah func runningApplication selesai di eksekusi
	defer logging()
	fmt.Println("Runnning Application")
}

func main() {

	runningApplication()
}
