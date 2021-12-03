package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type counter struct {
	t int
	f int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	counters := make(map[int]*counter)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := strconv.ParseInt(scanner.Text(), 2, 64)
		if err != nil {
			panic(err)
		}

		if size == 0 {
			size = len(scanner.Text())
		}

		for i := 0; i < size; i++ {
			c := counters[i]
			if c == nil {
				counters[i] = &counter{}
				c = counters[i]
			}

			if report&(1<<i) > 0 {
				c.t++
			} else {
				c.f++
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var gamma, mask int

	for i, c := range counters {
		if c.t > c.f {
			gamma += 1 << i
		}
		mask += 1 << i
	}

	epsilon := gamma ^ mask
	fmt.Printf("gamma: %d, epsilon: %d, answer: %d\n", gamma, epsilon, gamma*epsilon)
}
