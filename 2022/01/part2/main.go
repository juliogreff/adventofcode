package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		elf    int
		totals []int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totals = append(totals, int(elf))
			elf = 0
			continue
		}

		elf += mustparse.Int(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(totals)

	var answer int
	answer += totals[len(totals)-1]
	answer += totals[len(totals)-2]
	answer += totals[len(totals)-3]

	fmt.Printf("%d\n", answer)
}
