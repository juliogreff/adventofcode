package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	measurement := math.MaxInt
	increases := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if m > measurement {
			increases++
		}

		measurement = m
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", increases)
}
