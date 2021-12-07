package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const StartSize = 5000

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var maxPos int
	crabs := make([]int, StartSize, StartSize)
	scanner := bufio.NewScanner(file)
	if ok := scanner.Scan(); !ok {
		panic("nothing to read!")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, str := range strings.Split(scanner.Text(), ",") {
		c, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		crabs[c]++

		if c > maxPos {
			maxPos = c
		}
	}

	crabs = crabs[:maxPos+1]

	minFuel := math.MaxInt
	for pos := range crabs {
		f := countFuel(crabs, pos)
		if f < minFuel {
			minFuel = f
		}
	}

	fmt.Printf("%d\n", minFuel)
}

func countFuel(crabs []int, pos int) int {
	var f int

	for i, c := range crabs {
		abs := int(math.Abs(float64(pos - i)))
		f += c * abs
	}

	return f
}
