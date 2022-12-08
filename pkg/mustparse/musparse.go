package mustparse

import (
	"strconv"
	"strings"
)

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

func Ints(str string) []int {
	ints := make([]int, 0, len(str))

	for _, c := range str {
		ints = append(ints, Int(string(c)))
	}

	return ints
}

func SplitInts(str, sep string) []int {
	separated := strings.Split(str, sep)
	ints := make([]int, 0, len(separated))

	for _, i := range separated {
		ints = append(ints, Int(i))
	}

	return ints
}

func Float(str string) float64 {
	f, err := strconv.ParseFloat(str, bits)
	if err != nil {
		panic(err)
	}

	return f
}
