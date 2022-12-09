package xy

import "github.com/juliogreff/adventofcode/pkg/intmath"

type XY struct {
	X int
	Y int
}

func (a XY) Distance(b XY) XY {
	return XY{
		X: b.X - a.X,
		Y: b.Y - a.Y,
	}
}

func (a XY) ManhattanDistance(b XY) int {
	return intmath.Abs(a.X-b.X) + intmath.Abs(a.Y-b.Y)
}
