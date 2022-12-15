package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/intmath"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

type segment struct {
	min int
	max int
}

func main() {
	depth := 2000000

	file := os.Args[1]
	if strings.Contains(file, "test") {
		depth = 10
	}

	mustread.File(file, func(scanner *bufio.Scanner) {
		var segments []segment

		for scanner.Scan() {
			line := scanner.Text()

			var x, y, bx, by int

			fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bx, &by)

			sensor := xy.XY{x, y}
			beacon := xy.XY{bx, by}

			manhattan := beacon.ManhattanDistance(sensor)

			if depth >= (sensor.Y-manhattan) && depth <= (sensor.Y+manhattan) {
				s := segment{
					sensor.X - manhattan + intmath.Abs(depth-sensor.Y),
					sensor.X + manhattan - intmath.Abs(depth-sensor.Y),
				}

				segments = append(segments, s)
			}
		}

		line := make(map[int]struct{})
		for _, s := range segments {
			for i := s.min; i <= s.max; i++ {
				line[i] = struct{}{}
			}
		}

		fmt.Printf("%d\n", len(line)-1)
	})
}
