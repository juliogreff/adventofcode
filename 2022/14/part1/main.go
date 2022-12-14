package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/intmath"
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
		minX := math.MaxInt
		maxX := math.MinInt
		maxY := math.MinInt

		for scanner.Scan() {
			var from *xy.XY

			for _, l := range strings.Split(scanner.Text(), " -> ") {
				x := mustparse.XY(l)

				minX = intmath.Min(x.X, minX)
				maxX = intmath.Max(x.X, maxX)
				maxY = intmath.Max(x.Y, maxY)

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

		source := xy.XY{500, 0}
		for {
			next, ok := findNext(cave, source, minX, maxX, maxY)
			if !ok {
				break
			}

			cave[next] = struct{}{}
			answer++
		}

		fmt.Printf("%d\n", answer)

	})
}

func findNext(cave map[xy.XY]struct{}, start xy.XY, minX, maxX, maxY int) (xy.XY, bool) {
	next := start
	for {
		if next.X < minX || next.X > maxX || next.Y > maxY {
			return next, false
		}

		if _, ok := cave[xy.XY{next.X, next.Y + 1}]; !ok {
			next = xy.XY{next.X, next.Y + 1}
		} else if _, ok := cave[xy.XY{next.X - 1, next.Y + 1}]; !ok {
			next = xy.XY{next.X - 1, next.Y + 1}
		} else if _, ok := cave[xy.XY{next.X + 1, next.Y + 1}]; !ok {
			next = xy.XY{next.X + 1, next.Y + 1}
		} else {
			return next, true
		}
	}

	return next, false
}
