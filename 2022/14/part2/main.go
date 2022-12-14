package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			answer int
		)

		cave := make(map[xy.XY]struct{})
		maxY := math.MinInt

		for scanner.Scan() {
			var from *xy.XY

			for _, l := range strings.Split(scanner.Text(), " -> ") {
				x := mustparse.XY(l)

				if x.Y > maxY {
					maxY = x.Y
				}

				if from == nil {
					from = &x
					continue
				}

				from.StraightWalk(x, func(dst xy.XY) {
					cave[dst] = struct{}{}
				})

				from = &x
			}
		}

		maxY += 2

		source := xy.XY{500, 0}
		for {
			next, ok := findNext(cave, source, maxY)

			cave[next] = struct{}{}
			answer++

			if !ok {
				break
			}
		}

		fmt.Printf("%d\n", answer)

	})
}

func findNext(cave map[xy.XY]struct{}, start xy.XY, maxY int) (xy.XY, bool) {
	next := start

	for {
		cave[xy.XY{next.X, maxY}] = struct{}{}
		cave[xy.XY{next.X + 1, maxY}] = struct{}{}
		cave[xy.XY{next.X - 1, maxY}] = struct{}{}

		down := xy.XY{next.X, next.Y + 1}
		left := xy.XY{next.X - 1, next.Y + 1}
		right := xy.XY{next.X + 1, next.Y + 1}

		if _, ok := cave[down]; !ok {
			next = down
		} else if _, ok := cave[left]; !ok {
			next = left
		} else if _, ok := cave[right]; !ok {
			next = right
		} else {
			return next, next.Y != 0
		}
	}

	return next, false
}
