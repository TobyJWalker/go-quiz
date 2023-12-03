package main

import (
	"flag"
)

// create flags
var (
	timer = flag.Int("t", 30, "Time limit for the quiz in seconds")
	csv_path = flag.String("csv", "problems.csv", "Path to the csv file")
)