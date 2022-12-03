package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

var winsAgainst = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var losesAgainst = map[string]string{
	"C": "A",
	"A": "B",
	"B": "C",
}

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var answer int
		for scanner.Scan() {
			line := scanner.Text()

			round := strings.Split(line, " ")

			var (
				scoreMatch int
				chosen     string
			)

			switch round[1] {
			case "X":
				scoreMatch = 0
				chosen = winsAgainst[round[0]]
			case "Y":
				scoreMatch = 3
				chosen = round[0]
			case "Z":
				scoreMatch = 6
				chosen = losesAgainst[round[0]]
			}

			scoreChosen := scores[chosen]

			answer += scoreChosen + scoreMatch
		}

		fmt.Printf("%d\n", answer)
	})
}
