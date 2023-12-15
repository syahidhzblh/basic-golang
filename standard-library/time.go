package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2023, time.May, 17, 17, 10, 10, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "02-01-2006 15:04:05"

	value := "17-05-1999 05:00:00"

	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}
}
