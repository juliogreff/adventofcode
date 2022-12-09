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
		head := &xy.XY{}
		tail := &xy.XY{}

		visits := map[xy.XY]struct{}{
			*tail: struct{}{},
		}

		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Split(line, " ")
			amount := mustparse.Int(parts[1])

			for i := 0; i < amount; i++ {
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

				distance := tail.Distance(*head)
				mDistance := tail.ManhattanDistance(*head)

				if mDistance > 2 {
					tail.X += intmath.Sign(distance.X)
					tail.Y += intmath.Sign(distance.Y)
				} else {
					if intmath.Abs(distance.X) > 1 {
						tail.X += intmath.Sign(distance.X)
					}

					if intmath.Abs(distance.Y) > 1 {
						tail.Y += intmath.Sign(distance.Y)
					}
				}

				visits[*tail] = struct{}{}
			}
		}

		fmt.Printf("%d\n", len(visits))
	})
}
