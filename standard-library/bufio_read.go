package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	input := strings.NewReader("ini baris satu\nini baris dua\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
