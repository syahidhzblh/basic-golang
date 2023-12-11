package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error,", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error Code : 404")
	}
}

func main() {
	runApp(true)
	fmt.Println("Syahid Hizbullah")
}
