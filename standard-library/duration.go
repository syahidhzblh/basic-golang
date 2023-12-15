package main

import (
	"fmt"
	"time"
)

func main() {

	duration1 := time.Second * 100
	duration2 := time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration %d", duration3)
}
