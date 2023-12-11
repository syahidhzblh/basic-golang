package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 5
	var d = a + b*c

	fmt.Println(d)

	var i = 10
	i += 10 // i + 10
	fmt.Println(i)

	i++ // i + 1
	fmt.Println(i)

	i-- // i - 1
	fmt.Println(i)
}
