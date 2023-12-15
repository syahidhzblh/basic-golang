package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Syahid Hizbullah", "Syahid"))
	fmt.Println(strings.Split("Syahid Hizbullah", " "))
	fmt.Println(strings.ToLower("Syahid Hizbullah"))
	fmt.Println(strings.ToUpper("Syahid Hizbullah"))
	fmt.Println(strings.Trim("     Syahid Hizbullah      ", " "))
	fmt.Println(strings.ReplaceAll("Syahid Hizbullah Syahid", "Syahid", "Hisbul"))
}
