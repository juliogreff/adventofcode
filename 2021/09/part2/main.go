package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type heightmap [][]int

func newHeightmap() heightmap {
	return make([][]int, 0, 10)
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int
	hmap := newHeightmap()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		positions := strings.Split(line, "")
		row := make([]int, 0, len(positions))
		for _, p := range positions {
			i, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}

			row = append(row, i)
		}

		hmap = append(hmap, row)
	}

	var basins []int
	for i, row := range hmap {
		for j, n := range row {
			if isLowPoint(hmap, i, j) {
				c := 1 + countFlowDownwards(hmap, i, j, n)
				basins = append(basins, c)
			}
		}
	}

	sort.Ints(basins)

	largest := basins[len(basins)-3:]
	for _, b := range largest {
		if answer == 0 {
			answer = b
		} else {
			answer *= b
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}

func countFlowDownwards(hmap heightmap, i, j, n int) int {
	var sum int

	fn := func(i int) bool {
		return i < 9 && i > n
	}

	// up
	if i > 0 && fn(hmap[i-1][j]) {
		sum += 1 + countFlowDownwards(hmap, i-1, j, hmap[i-1][j])
		hmap[i-1][j] = 9
	}

	// down
	if i < len(hmap)-1 && fn(hmap[i+1][j]) {
		sum += 1 + countFlowDownwards(hmap, i+1, j, hmap[i+1][j])
		hmap[i+1][j] = 9
	}

	// left
	if j > 0 && fn(hmap[i][j-1]) {
		sum += 1 + countFlowDownwards(hmap, i, j-1, hmap[i][j-1])
		hmap[i][j-1] = 9
	}

	// right
	if j < len(hmap[i])-1 && fn(hmap[i][j+1]) {
		sum += 1 + countFlowDownwards(hmap, i, j+1, hmap[i][j+1])
		hmap[i][j+1] = 9
	}

	return sum
}

func isLowPoint(hmap heightmap, i, j int) bool {
	n := hmap[i][j]
	// up
	if i > 0 && hmap[i-1][j] <= n {
		return false
	}

	// down
	if i < len(hmap)-1 && hmap[i+1][j] <= n {
		return false
	}

	// left
	if j > 0 && hmap[i][j-1] <= n {
		return false
	}

	// right
	if j < len(hmap[i])-1 && hmap[i][j+1] <= n {
		return false
	}

	return true
}
