package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var winsAgainst = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		round := strings.Split(line, " ")

		switch round[1] {
		case "X":
			round[1] = "A"
		case "Y":
			round[1] = "B"
		case "Z":
			round[1] = "C"
		}

		scoreChosen := scores[round[1]]
		var scoreMatch int

		if round[0] == round[1] {
			scoreMatch = 3
		} else if winsAgainst[round[1]] == round[0] {
			scoreMatch = 6
		}

		// fmt.Printf("%s: %d %d\n", round, scoreChosen, scoreMatch)

		answer += scoreChosen + scoreMatch
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}
