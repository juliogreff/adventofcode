package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	WindowSize = 3
	RingSize   = WindowSize + 1
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ring := make([]int, RingSize, RingSize)
	iter := 0
	increases := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		idx := iter % (RingSize)
		ring[idx] = m

		iter++

		if iter > WindowSize {
			sumBefore := calculateSum(ring, iter-1)
			sumAfter := calculateSum(ring, iter)

			if sumAfter > sumBefore {
				increases++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", increases)
}

func calculateSum(ring []int, iter int) int {
	sum := 0
	for i := 0; i < WindowSize; i++ {
		idx := (iter - i - 1) % (RingSize)
		sum += ring[idx]
	}

	return sum
}
