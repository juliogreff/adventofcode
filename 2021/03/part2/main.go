package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type counter struct {
	t int
	f int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		size    int
		reports []int
	)

	counters := make(map[int]*counter)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := strconv.ParseInt(scanner.Text(), 2, 64)
		if err != nil {
			panic(err)
		}

		reports = append(reports, int(report))

		if size == 0 {
			size = len(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	o2, co2 := filterReports(counters, size-1, reports, reports)

	fmt.Printf("o2: %d, co2: %d, answer: %d\n", o2, co2, o2*co2)
}

func buildCounter(i int, reports []int) counter {
	var c counter
	for _, report := range reports {
		if isOne(report, i) {
			c.t++
		} else {
			c.f++
		}
	}

	return c
}

func filterReports(counters map[int]*counter, i int, o2, co2 []int) (int, int) {
	if i == 0 {
		return o2[0], co2[0]
	}

	if len(o2) > 1 {
		counter := buildCounter(i, o2)
		o2 = selekt(o2, func(report int) bool {
			return selectO2(counter, i, report)
		})
	}

	if len(co2) > 1 {
		counter := buildCounter(i, co2)
		co2 = selekt(co2, func(report int) bool {
			return selectCO2(counter, i, report)
		})
	}

	return filterReports(counters, i-1, o2, co2)
}

func selekt(s []int, fn func(int) bool) []int {
	selected := make([]int, 0, len(s))
	for _, v := range s {
		if fn(v) {
			selected = append(selected, v)
		}
	}

	return selected
}

func selectO2(counter counter, i int, report int) bool {
	if counter.t >= counter.f {
		return isOne(report, i)
	} else {
		return isZero(report, i)
	}
}

func selectCO2(counter counter, i int, report int) bool {
	if counter.f <= counter.t {
		return isZero(report, i)
	} else {
		return isOne(report, i)
	}
}

func isOne(v, i int) bool {
	return v&(1<<i) > 0
}

func isZero(v, i int) bool {
	return v&(1<<i) == 0
}
