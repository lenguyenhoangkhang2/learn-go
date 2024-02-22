package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readFile() {
	file, err := os.Open("./tmp/MOCK_DATA.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}

func writeFile() {
	file2, err := os.Create("./tmp/data.csv")
	if err != nil {
		panic(err)
	}

	defer file2.Close()

	writer := csv.NewWriter(file2)

	defer writer.Flush()

	headers := []string{"name", "age", "gender"}

	data1 := [][]string{
		{"Alice", "25", "Female"},
		{"Bob", "30", "Male"},
		{"Charlie", "35", "Male"},
	}

	writer.Write(headers)

	for _, row := range data1 {
		writer.Write(row)
	}
}

func main() {
    readFile()
}
