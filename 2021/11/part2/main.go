package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Steps = 100

type pos struct {
	x int
	y int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int
	octopi := make([][]int, 0, 10)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, "")
		row := make([]int, 0, len(levels))
		for _, p := range levels {
			i, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}

			row = append(row, i)
		}

		octopi = append(octopi, row)
	}

	for {
		var flashes int
		flashed := make(map[pos]struct{})
		for i := range octopi {
			for j := range octopi[i] {
				flashes += flash(octopi, flashed, i, j)
			}
		}

		answer++

		if flashes == 100 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}

func flash(octopi [][]int, flashed map[pos]struct{}, i, j int) int {
	var flashes int

	if i < 0 || i >= len(octopi) {
		return flashes
	}

	if j < 0 || j >= len(octopi[i]) {
		return flashes
	}

	if _, ok := flashed[pos{i, j}]; ok {
		// fmt.Printf("%d %d already flashed\n", i, j)
		return flashes
	}

	octopi[i][j]++

	if octopi[i][j] > 9 {
		// fmt.Printf("%d %d flashed\n", i, j)

		flashed[pos{i, j}] = struct{}{}
		octopi[i][j] = 0
		flashes += 1
		flashes += flash(octopi, flashed, i+1, j)   // up
		flashes += flash(octopi, flashed, i+1, j+1) // up+right
		flashes += flash(octopi, flashed, i+1, j-1) // up+left
		flashes += flash(octopi, flashed, i-1, j)   // down
		flashes += flash(octopi, flashed, i-1, j+1) // down+right
		flashes += flash(octopi, flashed, i-1, j-1) // down+left
		flashes += flash(octopi, flashed, i, j+1)   // right
		flashes += flash(octopi, flashed, i, j-1)   // left
	}

	return flashes
}
