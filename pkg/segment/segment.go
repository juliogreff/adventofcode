package segment

import (
	"fmt"

	"github.com/juliogreff/adventofcode/pkg/intmath"
)

type Segment struct {
	Min int
	Max int
}

func (s Segment) Clamp(min, max int) Segment {
	return Segment{
		Min: intmath.Max(s.Min, min),
		Max: intmath.Min(s.Max, max),
	}
}

func (a Segment) Overlap(b Segment) bool {
	return intmath.Max(a.Min, a.Max) >= intmath.Min(b.Min, b.Max) &&
		intmath.Max(b.Min, b.Max) >= intmath.Min(a.Min, a.Max)
}

func (a Segment) Expand(b Segment) Segment {
	if !a.Overlap(b) {
		panic(fmt.Errorf("segment %v does not overlap with %v, cannot expand", a, b))
	}

	return Segment{
		Min: intmath.Min(a.Min, b.Min),
		Max: intmath.Max(a.Max, b.Max),
	}
}
