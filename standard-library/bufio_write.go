package main

import (
	"bufio"
	"os"
)

func main() {

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Syahid Hizbullah\n")
	_, _ = writer.WriteString("Hisbul Hisbul")

	writer.Flush()
}
