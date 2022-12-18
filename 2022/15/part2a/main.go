package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliogreff/adventofcode/pkg/intmath"
	"github.com/juliogreff/adventofcode/pkg/lists"
	"github.com/juliogreff/adventofcode/pkg/mustread"
	"github.com/juliogreff/adventofcode/pkg/segment"
	"github.com/juliogreff/adventofcode/pkg/xy"
)

type signal struct {
	sensor   xy.XY
	distance int
}

func main() {
	max := 4_000_000

	file := os.Args[1]
	if strings.Contains(file, "test") {
		max = 20
	}

	mustread.File(file, func(scanner *bufio.Scanner) {
		var signals []signal

		for scanner.Scan() {
			var x, y, bx, by int

			fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bx, &by)

			sensor := xy.XY{x, y}
			beacon := xy.XY{bx, by}

			signals = append(signals, signal{sensor, sensor.ManhattanDistance(beacon)})
		}

		segmentsByDepth := make([][]segment.Segment, max, max)
		for depth := 0; depth < max; depth++ {
			segmentsByDepth[depth] = make([]segment.Segment, 0, 16)

			signals := selectSignals(signals, depth)
			for _, signal := range signals {
				sensor := signal.sensor
				distance := signal.distance

				s := segment.Segment{
					intmath.Max(0, sensor.X-distance+intmath.Abs(depth-sensor.Y)),
					intmath.Min(max, sensor.X+distance-intmath.Abs(depth-sensor.Y)),
				}

				segmentsByDepth[depth] = append(segmentsByDepth[depth], s)
			}
		}

		beacon := findGap(segmentsByDepth, max)
		answer := beacon.X*4000000 + beacon.Y
		fmt.Printf("%d\n", answer)
	})
}

func selectSignals(signals []signal, depth int) []signal {
	selected := make([]signal, 0, len(signals))

	for _, signal := range signals {
		sensor := signal.sensor
		distance := signal.distance

		if depth >= (sensor.Y-distance) && depth <= (sensor.Y+distance) {
			selected = append(selected, signal)
		}
	}

	return selected
}

func findGap(segmentsByDepth [][]segment.Segment, max int) xy.XY {
	for depth, segments := range segmentsByDepth {
		lists.Sort(segments, func(a, b segment.Segment) bool {
			return a.Min < b.Min
		})

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
