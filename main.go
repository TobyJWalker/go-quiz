package main

import (
	"bufio"
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

// create input reader
var io = bufio.NewReader(os.Stdin)

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

	// create channels
	timer_ch := make(chan interface{})
	quiz_ch := make(chan interface{})
	done_ch := make(chan interface{})

	// prompt user to start
	fmt.Println("Press enter to start the quiz")
	io.ReadString('\n')

	// initailise correct answers
	correct := 0

	// function to listen for correct answers
	go func() {
		for {
			select {
			case <-quiz_ch:
				correct++
			case <-timer_ch:
				done_ch <- true
			}
		}
	}()

	// start quiz
	go startQuiz(records, timer_ch, quiz_ch, done_ch)

	

}

// function to start quiz
func startQuiz(problems [][]string, timer_ch, quiz_ch, done_ch chan interface{}) {
	
	// track index
	index := 0

	// loop for asking questions
	for {
		select {
			// handle timeout
			case <-timer_ch:
				fmt.Println("Time's up!")
				return
			
			// ask questions
			default:
				// display question
				fmt.Printf("Question %d: %s = ", index+1, problems[index][0])

				// read answer
				answer, err := io.ReadString('\n')
				if err != nil {
					panic("Error reading input")
				}

				// check answer
				if answer == problems[index][1] {
					quiz_ch <- true // signal correct answer
				}

				// increment index
				index++
		}
	}

}

