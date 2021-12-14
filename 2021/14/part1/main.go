package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const Steps = 10

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var polymer []string
	rules := make(map[string]map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		polymer = strings.Split(line, "")
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		pair := strings.Split(parts[0], "")
		insertion := parts[1]

		if _, ok := rules[pair[0]]; !ok {
			rules[pair[0]] = make(map[string]string)
		}

		rules[pair[0]][pair[1]] = insertion
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for step := 0; step < Steps; step++ {
		for i := 0; i < len(polymer)-1; i++ {
			_, ok := rules[polymer[i]]
			if !ok {
				continue
			}

			insert, ok := rules[polymer[i]][polymer[i+1]]
			if !ok {
				continue
			}

			polymer = append(polymer[:i+1], append([]string{insert}, polymer[i+1:]...)...)
			i++
		}
	}

	min := math.MaxInt
	max := 0

	counters := make(map[string]int)
	for _, p := range polymer {
		counters[p]++
	}

	for _, c := range counters {
		if c < min {
			min = c
		} else if c > max {
			max = c
		}
	}

	fmt.Printf("%d\n", max-min)
}
