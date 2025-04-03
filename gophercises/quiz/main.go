package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// go build main.go && ./main --help
// ./main -csv="problems.csv"
func main() {
	// terminal setup
	csvFileNamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	// read the file
	file, err := os.Open(*csvFileNamePtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFileNamePtr))
	}
	// parse the csv format
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	// => struct problem [decouple from how they came into our program]
	problems := parseLines(lines)
	// keep track of score
	correct := 0
	for i, p := range problems {
		// print out
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// get response, automatically trimming space
		var answer string
		fmt.Scanf("%s\n", &answer)
		// check
		if answer == p.a {
			correct++
			//fmt.Println("Correct!")
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	// avoid resizing of append
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return ret
}

type problem struct {
	q string
	a string
}

// reusable function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
