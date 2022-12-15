package segment

import "github.com/juliogreff/adventofcode/pkg/intmath"

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

// func overlap(a, b segment.) bool {
// 	return intmath.Max(a.min, a.max) >= intmath.Min(b.min, b.max) && intmath.Max(b.min, b.max) >= intmath.Min(a.min, a.max)
// }
//
// func expand(a, b segment) segment {
// 	return segment{
// 		min: intmath.Min(a.min, b.min),
// 		max: intmath.Max(a.max, b.max),
// 	}
// }
