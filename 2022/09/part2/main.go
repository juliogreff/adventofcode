package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/intmath"
	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		knots := []*xy.XY{
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
			&xy.XY{},
		}

		visits := map[xy.XY]struct{}{
			*knots[9]: struct{}{},
		}

		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Split(line, " ")
			amount := mustparse.Int(parts[1])

			for i := 0; i < amount; i++ {
				head := knots[0]
				tail := knots[9]

				switch parts[0] {
				case "R":
					head.X++
				case "L":
					head.X--
				case "U":
					head.Y++
				case "D":
					head.Y--
				}

				for i, knot := range knots[1:] {
					prev := knots[i]

					distance := knot.Distance(*prev)
					mDistance := knot.ManhattanDistance(*prev)

					if mDistance > 2 {
						knot.X += intmath.Sign(distance.X)
						knot.Y += intmath.Sign(distance.Y)
					} else {
						if intmath.Abs(distance.X) > 1 {
							knot.X += intmath.Sign(distance.X)
						}

						if intmath.Abs(distance.Y) > 1 {
							knot.Y += intmath.Sign(distance.Y)
						}
					}
				}

				visits[*tail] = struct{}{}
			}
		}

		fmt.Printf("%d\n", len(visits))
	})
}
