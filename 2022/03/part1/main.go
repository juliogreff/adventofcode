package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var answer int
		for scanner.Scan() {
			line := scanner.Text()
			set := make(map[byte]struct{})

			half := len(line) / 2

			for _, c := range []byte(line[:half]) {
				set[c] = struct{}{}
			}

			for _, c := range []byte(line[half:]) {
				_, ok := set[c]
				if ok {
					answer += prio(c)
					break
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
