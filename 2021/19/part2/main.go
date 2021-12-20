package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MinOverlap = 12

var orientations = []func(pos) pos{
	func(p pos) pos { return pos{p.x, p.y, p.z} },
	func(p pos) pos { return pos{-p.y, p.x, p.z} },
	func(p pos) pos { return pos{-p.x, -p.y, p.z} },
	func(p pos) pos { return pos{p.y, -p.x, p.z} },

	func(p pos) pos { return pos{-p.x, p.y, -p.z} },
	func(p pos) pos { return pos{p.y, p.x, -p.z} },
	func(p pos) pos { return pos{p.x, -p.y, -p.z} },
	func(p pos) pos { return pos{-p.y, -p.x, -p.z} },

	func(p pos) pos { return pos{-p.z, p.y, p.x} },
	func(p pos) pos { return pos{-p.z, p.x, -p.y} },
	func(p pos) pos { return pos{-p.z, -p.y, -p.x} },
	func(p pos) pos { return pos{-p.z, -p.x, p.y} },

	func(p pos) pos { return pos{p.z, p.y, -p.x} },
	func(p pos) pos { return pos{p.z, p.x, p.y} },
	func(p pos) pos { return pos{p.z, -p.y, p.x} },
	func(p pos) pos { return pos{p.z, -p.x, -p.y} },

	func(p pos) pos { return pos{p.x, -p.z, p.y} },
	func(p pos) pos { return pos{-p.y, -p.z, p.x} },
	func(p pos) pos { return pos{-p.x, -p.z, -p.y} },
	func(p pos) pos { return pos{p.y, -p.z, -p.x} },

	func(p pos) pos { return pos{p.x, p.z, -p.y} },
	func(p pos) pos { return pos{-p.y, p.z, -p.x} },
	func(p pos) pos { return pos{-p.x, p.z, p.y} },
	func(p pos) pos { return pos{p.y, p.z, p.x} },
}

type atob struct {
	a int
	b int
}

type pos struct {
	x int
	y int
	z int
}

func (p pos) ManhattanDelta(pp pos) float64 {
	return math.Abs(float64(p.x-pp.x)) + math.Abs(float64(p.y-pp.y)) + math.Abs(float64(p.z-pp.z))
}

func (p pos) Delta(pp pos) float64 {
	return math.Sqrt(
		math.Pow(float64(p.x-pp.x), 2) +
			math.Pow(float64(p.y-pp.y), 2) +
			math.Pow(float64(p.z-pp.z), 2),
	)
}

func (p pos) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (p pos) allOrientations() []pos {
	return []pos{}
}

type scanner []pos

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var scanners []scanner
	var s scanner

	txtScanner := bufio.NewScanner(file)
	for txtScanner.Scan() {
		line := txtScanner.Text()

		if line == "" {
			scanners = append(scanners, s)
			continue
		}

		if line[0:3] == "---" {
			s = scanner{}
			continue
		}

		s = append(s, parsePos(line))
	}

	if err := txtScanner.Err(); err != nil {
		panic(err)
	}

	positions := findAbsolutePositions(scanners)

	var max float64
	for _, a := range positions {
		for _, b := range positions {
			d := a.ManhattanDelta(b)
			if d > max {
				max = d
			}
		}
	}
	fmt.Printf("%f\n", max)
}

func findAbsolutePositions(scanners []scanner) map[int]pos {
	reoriented := map[int]pos{
		0: pos{0, 0, 0},
	}

	for len(reoriented) < len(scanners) {
		for i, _ := range reoriented {
			a := scanners[i]
			for j, b := range scanners {
				_, ok := reoriented[j]
				if i == j || ok {
					continue
				}

				relPos, o, _ := findRelativePosition(a, b)
				if o == nil {
					continue
				}

				b = reorientScanner(b, o, relPos)
				scanners[j] = b
				reoriented[j] = pos{-relPos.x, -relPos.y, -relPos.z}
			}
		}
	}

	return reoriented
}

func reorientBeacon(b pos, o func(pos) pos, relPos pos) pos {
	b = o(b)
	return pos{
		b.x - relPos.x,
		b.y - relPos.y,
		b.z - relPos.z,
	}
}

func reorientScanner(o scanner, orientation func(pos) pos, relPos pos) scanner {
	var n scanner

	for _, beacon := range o {
		n = append(n, reorientBeacon(beacon, orientation, relPos))
	}

	return n
}

func findRelativePosition(a, b scanner) (pos, func(pos) pos, float64) {
	for _, o := range orientations {
		deltas := make(map[float64]int)

		for _, src := range a {
			for _, dst := range b {
				translatedDst := o(dst)

				delta := src.Delta(translatedDst)

				deltas[delta]++

				if deltas[delta] >= MinOverlap {
					return pos{
						translatedDst.x - src.x,
						translatedDst.y - src.y,
						translatedDst.z - src.z,
					}, o, delta
				}
			}
		}
	}

	return pos{}, nil, 0
}

func parsePos(line string) pos {
	parts := strings.Split(line, ",")
	return pos{
		x: parseInt(parts[0]),
		y: parseInt(parts[1]),
		z: parseInt(parts[2]),
	}
}

func parseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return n
}
