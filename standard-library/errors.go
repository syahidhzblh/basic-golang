package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

type result struct {
	name string
}

func (v *result) Error() string {
	return v.name
}

func GetById(id string) error {

	if id == "" {
		return ValidationError
	}
	if id != "syahid" {
		return NotFoundError
	}

	return &result{name: id}
}

func main() {

	data := GetById("syahid")

	if data != nil {
		if errors.Is(data, ValidationError) {
			fmt.Println("validation error:")
		} else if errors.Is(data, NotFoundError) {
			fmt.Println("not found error:")
		} else {
			if r, ok := data.(*result); ok {
				fmt.Printf("Success, Hai user %s", r.name)
			} else {
				fmt.Println("unknown error")
			}
		}
	}
}
