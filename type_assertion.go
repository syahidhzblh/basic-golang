package main

import "fmt"

func random() any {
	return "Syahid"
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type", value)
	}
}
