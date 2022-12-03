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
			fmt.Printf("%s\n", line)
		}

		fmt.Printf("%d\n", answer)
	})
}
