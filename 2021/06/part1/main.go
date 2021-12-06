package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxLifespan = 9
	Offset      = 2
	Days        = 80
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cohorts = make([]int, MaxLifespan, MaxLifespan)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		for _, part := range parts {
			d, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			cohorts[d]++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var answer int
	for i := 0; i <= Days; i++ {
		var carry int
		answer = 0
		for d := MaxLifespan - 1; d >= 0; d-- {
			answer += carry
			cohorts[d], carry = carry, cohorts[d]

			if d == 0 {
				answer += carry
				cohorts[MaxLifespan-1] = carry
				cohorts[MaxLifespan-1-Offset] += carry
			}
		}
	}

	fmt.Printf("%d\n", answer)
}
