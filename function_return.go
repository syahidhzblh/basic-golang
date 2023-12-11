package main

import "fmt"

func getHello(name string) string {

	hello := "Hello " + name
	return hello
}

func main() {

	//Simpan dalam variable result then print
	result := getHello("Syahid Hizbullah")
	fmt.Println(result)

	//Langsung eksekusi by calling func getHello on Println
	fmt.Println(getHello("Syahid"))
}
