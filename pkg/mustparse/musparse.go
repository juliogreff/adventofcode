package mustparse

import "strconv"

const (
	decimal = 10
	bits    = 64
)

func Int(str string) int {
	i, err := strconv.ParseInt(str, decimal, bits)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func Float(str string) float64 {
	f, err := strconv.ParseFloat(str, bits)
	if err != nil {
		panic(err)
	}

	return f
}
