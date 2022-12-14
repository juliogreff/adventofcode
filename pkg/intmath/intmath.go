package intmath

import "math"

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func Sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}

	return 0
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
