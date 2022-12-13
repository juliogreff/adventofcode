package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			answer int
			lines  [][]any
		)

		for scanner.Scan() {
			line := scanner.Text()

			if line == "" {
				continue
			}

			var parsed []any
			err := json.Unmarshal([]byte(line), &parsed)
			if err != nil {
				panic(err)
			}

			lines = append(lines, parsed)
		}

		for i := 0; i < len(lines); i += 2 {
			left := lines[i]
			right := lines[i+1]

			c, _ := compare(left, right)

			if c {
				answer += i/2 + 1
			}
		}

		fmt.Printf("%d\n", answer)
	})
}

func compare(left any, right any) (bool, bool) {
	lInt, lIntOk := left.(float64)
	rInt, rIntOk := right.(float64)

	if lIntOk && rIntOk {
		return lInt <= rInt, lInt == rInt
	}

	var lLst, rLst []any

	if lIntOk {
		lLst = []any{lInt}
	} else {
		lLst = left.([]any)
	}

	if rIntOk {
		rLst = []any{rInt}
	} else {
		rLst = right.([]any)
	}

	for i := range lLst {
		if i >= len(rLst) {
			return false, false
		}

		ok, c := compare(lLst[i], rLst[i])
		if c {
			continue
		}

		return ok, false
	}

	return true, len(lLst) == len(rLst)
}
