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

func (a XY) StraightWalk(b XY, fn func(XY)) {
	if a.X == b.X {
		for i := intmath.Min(a.Y, b.Y); i <= intmath.Max(a.Y, b.Y); i++ {
			fn(XY{a.X, i})
		}
	} else if a.Y == b.Y {
		for i := intmath.Min(a.X, b.X); i <= intmath.Max(a.X, b.X); i++ {
			fn(XY{i, a.Y})
		}
	} else {
		panic("diagonal path")
	}
}
