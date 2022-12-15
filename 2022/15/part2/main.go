package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/intmath"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/segment"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

type sortable []segment.Segment

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Swap(i, j int) {
	a := s[i]
	b := s[j]

	s[i] = b
	s[j] = a
}

func (s sortable) Less(i, j int) bool {
	return s[i].Min < s[j].Min
}

func main() {
	max := 4000000

	file := os.Args[1]
	if strings.Contains(file, "test") {
		max = 20
	}

	mustread.File(file, func(scanner *bufio.Scanner) {
		segmentsByDepth := make([]sortable, max, max)

		for scanner.Scan() {
			line := scanner.Text()

			var x, y, bx, by int

			fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bx, &by)

			sensor := xy.XY{x, y}
			beacon := xy.XY{bx, by}

			manhattan := beacon.ManhattanDistance(sensor)

			var segments int

			for depth := intmath.Max(0, sensor.Y-manhattan); depth < intmath.Min(max, sensor.Y+manhattan); depth++ {
				s := segment.Segment{
					sensor.X - manhattan + intmath.Abs(depth-sensor.Y),
					sensor.X + manhattan - intmath.Abs(depth-sensor.Y),
				}.Clamp(0, max)

				segmentsByDepth[depth] = append(segmentsByDepth[depth], s)

				segments++
			}
		}

		beacon := findGap(segmentsByDepth, max)
		answer := beacon.X*4000000 + beacon.Y
		fmt.Printf("%d\n", answer)
	})
}

func findGap(segmentsByDepth []sortable, max int) xy.XY {
	for depth, segments := range segmentsByDepth {
		sort.Sort(segments)

		gap := segment.Segment{
			Min: 0,
			Max: max,
		}

		for _, s := range segments {
			if s.Max <= gap.Min || s.Min >= gap.Max {
				// segment ends before start of the gap, or starts before the end of the gap
				continue
			}

			if s.Min <= gap.Min && s.Max >= gap.Max {
				gap.Min = gap.Max
				break
			}

			if s.Min <= gap.Min && s.Max < gap.Max {
				gap.Min = s.Max
			} else if s.Min >= gap.Min && s.Min < gap.Max {
				gap.Max = s.Min
			}
		}

		// there's exactly a one-space gap
		if gap.Max-gap.Min == 2 {
			return xy.XY{gap.Min + 1, depth}
		}
	}

	panic("unreachable")
}