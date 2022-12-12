package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/juliogreff/adventofcode/pkg/graphs"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			answer int
			m      [][]int
			start  xy.XY
			end    xy.XY
		)

		for scanner.Scan() {
			line := scanner.Text()

			row := make([]int, 0, len(line))

			for i, c := range []byte(line) {
				if c == 'S' {
					start = xy.XY{i, len(m)}
					c = 'a'
				} else if c == 'E' {
					end = xy.XY{i, len(m)}
					c = 'z'
				}

				row = append(row, int(c-'a'))
			}

			m = append(m, row)
		}

		opts := graphs.DijkstraOpts{
			Cost: func(graph [][]int, src xy.XY, dst xy.XY) int {
				srcHeight := graph[src.Y][src.X]
				dstHeight := graph[dst.Y][dst.X]

				if dstHeight-srcHeight > 1 {
					return math.MaxInt
				}

				return 1
			},
		}

		answer = graphs.Dijkstra(m, start, end, opts)

		fmt.Printf("%d\n", answer)
	})
}
