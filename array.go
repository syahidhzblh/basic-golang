package main

import "fmt"

func main() {

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	var totals = [...]int{
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(totals)
	fmt.Println(len(totals))
	totals[0] = 100
	fmt.Println(totals)
}
