package main

import (
	"fmt"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustread"
)

func main() {
	trees := mustread.FileMapOfInts(os.Args[1])

	var answer int

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			d := distance(trees, i, j)

			if d > answer {
				answer = d
			}
		}
	}

	fmt.Printf("%d\n", answer)
}

func distance(trees [][]int, i, j int) int {
	rows := len(trees)
	columns := len(trees[0])

	distance := 1

	d := 0
	for x := i - 1; x >= 0; x-- {
		d++

		if trees[x][j] >= trees[i][j] {
			break
		}
	}
	distance *= d

	d = 0
	for x := i + 1; x < rows; x++ {
		d++

		if trees[x][j] >= trees[i][j] {
			break
		}
	}
	distance *= d

	d = 0
	for x := j - 1; x >= 0; x-- {
		d++

		if trees[i][x] >= trees[i][j] {
			break
		}
	}
	distance *= d

	d = 0
	for x := j + 1; x < columns; x++ {
		d++

		if trees[i][x] >= trees[i][j] {
			break
		}
	}
	distance *= d

	return distance
}
