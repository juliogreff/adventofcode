package main

import (
	"bufio"
	"fmt"
	"os"
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

	for i, row := range hmap {
		for j, n := range row {
			// up
			if i > 0 && hmap[i-1][j] <= n {
				continue
			}

			// down
			if i < len(hmap)-1 && hmap[i+1][j] <= n {
				continue
			}

			// left
			if j > 0 && hmap[i][j-1] <= n {
				continue
			}

			// right
			if j < len(row)-1 && hmap[i][j+1] <= n {
				continue
			}

			answer += 1 + n
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}
