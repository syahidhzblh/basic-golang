package main

import (
	"encoding/csv"
	"os"
)

func main() {

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"syahid", "hizbullah", "ganteng"})
	_ = writer.Write([]string{"hisbul", "ganteng", "ganteng"})
	_ = writer.Write([]string{"aku", "tetep", "ganteng"})

	writer.Flush()
}
