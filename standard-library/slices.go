package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"syahid", "hisbul", "bul", "idun", "zaid"}
	ages := []int{10, 20, 30, 40, 50}

	fmt.Println(slices.Min(ages))
	fmt.Println(slices.Max(ages))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "syahid"))
	fmt.Println(slices.Index(names, "idun"))
}
