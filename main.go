package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// create flags
var (
	timer = flag.Int("t", 30, "Time limit for the quiz in seconds")
	csv_path = flag.String("csv", "problems.csv", "Path to the csv file")
)

// create input reader
var io = bufio.NewScanner(os.Stdin)

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
	score_ch := make(chan int)

	// prompt user to start
	fmt.Println("Press enter to start the quiz")
	io.Scan()

	// start quiz
	go startQuiz(records, timer_ch, score_ch)

	// start timer func and close channels after
	go func() {
		time.Sleep(time.Duration(*timer) * time.Second)
		close(timer_ch)
	}()

	// get score
	correct := <-score_ch
	
	// display results and close score channel
	fmt.Printf("You scored %d out of %d\n", correct, len(records))
	close(score_ch)
}

// function to start quiz
func startQuiz(problems [][]string, timer_ch chan interface{}, score_ch chan int) {
	
	// init vars
	index := 0
	correct := 0

	// loop for asking questions
	for {
		select {
			// handle timeout
			case <-timer_ch:
				fmt.Println("Time's up!")
				score_ch <- correct
				return
			
			// ask questions
			default:
				// check if done
				if index >= len(problems) {
					fmt.Println("Quiz complete!")
					score_ch <- correct
					return
				}

				// display question
				fmt.Printf("Question %d: %s = ", index+1, problems[index][0])

				// read answer
				io.Scan()
				answer := io.Text()

				// format expected and given answer
				expected := strings.TrimSpace(strings.TrimRight(problems[index][1], "\n"))
				answer = strings.TrimSpace(strings.TrimRight(answer, "\n"))

				// check answer
				if answer == expected {
					correct++
				}

				// increment index
				index++
		}
	}

}

