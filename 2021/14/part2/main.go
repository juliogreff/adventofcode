package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const Steps = 40

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairs := make(map[string]int)
	counters := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		polymer := line
		for i := 0; i < len(polymer)-1; i++ {
			pairs[polymer[i:i+2]]++
			counters[string(polymer[i])]++
		}
		counters[string(polymer[len(polymer)-1])]++
	}

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for step := 0; step < Steps; step++ {
		newPairs := make(map[string]int)
		for k, v := range pairs {
			newPairs[k] = v
		}

		for pair := range pairs {
			if insert, ok := rules[pair]; ok {
				lhs := string(pair[0])
				rhs := string(pair[1])

				factor := pairs[pair]
				newPairs[pair] -= factor
				newPairs[lhs+insert] += factor
				newPairs[insert+rhs] += factor
				counters[insert] += factor
			}
		}

		pairs = newPairs
	}

	min := math.MaxInt
	max := 0
	for _, c := range counters {
		if c < min {
			min = c
		} else if c > max {
			max = c
		}
	}

	fmt.Printf("%d\n", max-min)
}
