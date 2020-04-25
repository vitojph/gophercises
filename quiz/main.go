package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return problems
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit (in seconds) to finish the quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Printf("Cannot open file: %s\n", *csvFileName)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Printf("Failed to parse de CSV file")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
	for i, problem := range problems {
		fmt.Printf("Question #%d: %s\n", i+1, problem.q)

		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou answered %d out of %d questions!. Your score is %.2f\n", counter, len(problems), float64(counter)/float64(len(problems)))
			return
		case answer := <-answerChan:
			if answer == problem.a {
				fmt.Println("Correct!")
				counter++
			} else {
				fmt.Println("Ooops!")
			}
		}
	}
	fmt.Printf("You answered %d out of %d questions!. Your score is %.2f\n", counter, len(problems), float64(counter)/float64(len(problems)))
}
