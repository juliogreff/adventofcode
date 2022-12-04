package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

type elf struct {
	min int
	max int
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var answer int
		for scanner.Scan() {
			line := scanner.Text()

			pairs := strings.Split(line, ",")

			ll := mustparse.SplitInts(pairs[0], "-")
			rr := mustparse.SplitInts(pairs[1], "-")

			var l, r elf

			l.min = ll[0]
			l.max = ll[1]
			r.min = rr[0]
			r.max = rr[1]

			if (l.min >= r.min && l.max <= r.max) || (r.min >= l.min && r.max <= l.max) {
				answer++
			}
		}

		fmt.Printf("%d\n", answer)
	})
}
