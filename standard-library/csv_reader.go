package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {

	csvString := "Syahid, Hizbullah, Ganteng\n" +
		"Hisbul, Hisbul, Ganteng\n" +
		"Aku, Tetep, Ganteng\n"

	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
