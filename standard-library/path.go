package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {

	//Windows using filepath
	fmt.Println(filepath.Dir("hello/main.go"))
	fmt.Println(filepath.Base("hello/main.go"))
	fmt.Println(filepath.Ext("hello/main.go"))
	fmt.Println(filepath.IsAbs("hello/main.go"))
	fmt.Println(filepath.IsLocal("hello/main.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go\n"))

	//Linux & Unix using filepath
	fmt.Println(path.Dir("hello/main.go"))
	fmt.Println(path.Base("hello/main.go"))
	fmt.Println(path.Ext("hello/main.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
