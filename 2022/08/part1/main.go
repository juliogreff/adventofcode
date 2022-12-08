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
			v := visible(trees, i, j)
			if v {
				answer++
			}
		}
	}

	answer += 2*len(trees) + 2*len(trees[0]) - 4

	fmt.Printf("%d\n", answer)
}

func visible(trees [][]int, i, j int) bool {
	rows := len(trees)
	columns := len(trees[0])

	visible := true
	for x := 0; x < i; x++ {
		if trees[x][j] >= trees[i][j] {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	for x := i + 1; x < rows; x++ {
		if trees[x][j] >= trees[i][j] {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	for x := 0; x < j; x++ {
		if trees[i][x] >= trees[i][j] {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	for x := j + 1; x < columns; x++ {
		if trees[i][x] >= trees[i][j] {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = false

	return false
}
