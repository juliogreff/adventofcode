package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cavern [][]int
	unvisited := make(map[pos]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var positions []int
		for x, p := range strings.Split(line, "") {
			i, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			positions = append(positions, i)

			unvisited[pos{x, len(cavern)}] = struct{}{}
		}

		cavern = append(cavern, positions)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	maxX := len(cavern[0]) - 1
	maxY := len(cavern) - 1

	node := pos{0, 0}
	delete(unvisited, node)

	distance := make(map[pos]int)
	distance[node] = 0

	visited := make(map[pos]struct{})
	dst := pos{maxX, maxY}

	for len(unvisited) > 0 {
		neighbors := []pos{
			{node.x, node.y + 1},
			{node.x, node.y - 1},
			{node.x + 1, node.y},
			{node.x - 1, node.y},
		}

		for _, n := range neighbors {
			if n.x < 0 || n.x > maxX || n.y < 0 || n.y > maxY {
				continue
			}

			if _, ok := visited[n]; ok {
				continue
			}

			cost := cavern[n.y][n.x]
			dtn := distance[node] + cost

			d, ok := distance[n]
			if !ok {
				d = math.MaxInt
			}

			if dtn < d {
				distance[n] = dtn
			}
		}

		visited[node] = struct{}{}
		if node == dst {
			break
		}

		node = getLowest(unvisited, distance)
	}

	fmt.Printf("%d\n", distance[dst])
}

func getLowest(unvisited map[pos]struct{}, distance map[pos]int) pos {
	var node pos
	var lowest = math.MaxInt

	for u := range unvisited {
		if d, ok := distance[u]; ok && d < lowest {
			node = u
			lowest = distance[u]
		}
	}

	delete(unvisited, node)

	return node
}
