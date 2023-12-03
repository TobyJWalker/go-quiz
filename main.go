package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// create flags
var (
	timer = flag.Int("t", 30, "Time limit for the quiz in seconds")
	csv_path = flag.String("csv", "problems.csv", "Path to the csv file")
)

func main() {

	// parse flags
	flag.Parse()

	// open csv file
	file, err := os.Open(*csv_path)
	if err != nil {
		panic("Error opening csv file")
	}
	defer file.Close()
	
	// create csv reader
	csv_reader := csv.NewReader(file)

	// attempt to read file
	records, err := csv_reader.ReadAll()

	// handle error
	if err != nil {
		panic("Error reading csv file")
	}

	fmt.Println(records)
}