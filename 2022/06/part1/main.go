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

			for i := 0; i <= len(line)-4; i++ {
				if allDifferent(line[i : i+4]) {
					answer = i + 4
					break
				}
			}
		}

		fmt.Printf("%d\n", answer)
	})
}

func allDifferent(str string) bool {
	set := make(map[byte]struct{})
	for i := 0; i < len(str); i++ {
		if _, ok := set[str[i]]; ok {
			return false
		}
		set[str[i]] = struct{}{}
	}

	return true
}
