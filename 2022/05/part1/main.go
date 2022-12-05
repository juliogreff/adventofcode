package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/juliogreff/adventofcode/pkg/lists"
	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			answer    string
			nStacks   int
			stacks    [][]string
			gotStacks bool
		)

		for scanner.Scan() {
			line := scanner.Text()

			if nStacks == 0 {
				nStacks = (len(line) + 1) / 4
				stacks = make([][]string, nStacks, nStacks)
			}

			if !gotStacks {
				if line[1] == '1' {
					gotStacks = true

					continue
				}

				for i := 0; i < nStacks; i++ {
					char := string(line[1+4*i])
					if char != " " {
						stacks[i] = append([]string{char}, stacks[i]...)
					}
				}

				continue
			}

			if line == "" {
				continue
			}

			r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
			matches := r.FindStringSubmatch(line)

			howMany := mustparse.Int(matches[1])
			from := mustparse.Int(matches[2]) - 1
			to := mustparse.Int(matches[3]) - 1

			idx := len(stacks[from]) - howMany
			pieces := stacks[from][idx:]
			stacks[from] = stacks[from][:idx]
			stacks[to] = append(stacks[to], lists.Reverse(pieces)...)
		}

		for _, s := range stacks {
			answer += s[len(s)-1]
		}

		fmt.Println(answer)
	})
}
