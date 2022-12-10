package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var answer int
		x := 1
		cycle := 1

		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")

			var n, cycles int

			switch parts[0] {
			case "addx":
				cycles = 2
				n = mustparse.Int(parts[1])
			case "noop":
				cycles = 1
			}

			for i := 0; i < cycles; i++ {
				if (cycle+20)%40 == 0 {
					answer += cycle * x
				}

				cycle++
			}

			x += n
		}

		fmt.Printf("%d\n", answer)
	})
}
