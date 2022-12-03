package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

type item struct {
	found   int
	lastElf int
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var answer int

	scan:
		for {
			set := make(map[byte]*item)

			for i := 1; i <= 3; i++ {
				ok := scanner.Scan()
				if !ok {
					break scan
				}

				line := scanner.Text()

				for _, c := range []byte(line) {
					it, ok := set[c]
					if !ok {
						set[c] = &item{}
						it = set[c]
					}

					if it.lastElf < i {
						it.lastElf = i
						it.found++
					}

					if set[c].found == 3 {
						answer += prio(c)
						break
					}
				}
			}
		}

		fmt.Printf("%d\n", answer)
	})
}

func prio(c byte) int {
	char := byte('A')
	offset := 27

	if c > 'Z' {
		char = 'a'
		offset = 1
	}

	return int(c-char) + offset
}
