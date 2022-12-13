package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

type sortable [][]any

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Swap(i, j int) {
	a := s[i]
	b := s[j]

	s[i] = b
	s[j] = a
}

func (s sortable) Less(i, j int) bool {
	c, _ := compare(s[i], s[j])
	return c
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var lines sortable
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

		d := []any{[]any{2.0}}
		dd := []any{[]any{6.0}}

		lines = append(lines, d, dd)

		sort.Sort(lines)

		answer := 1
		for i, l := range lines {
			if reflect.DeepEqual(l, d) || reflect.DeepEqual(l, dd) {
				answer *= i + 1
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
