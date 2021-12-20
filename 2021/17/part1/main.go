package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type pos struct {
	x int
	y int
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		minX := parseInt(matches[1])
		maxX := parseInt(matches[2])
		minY := parseInt(matches[3])
		maxY := parseInt(matches[4])

		minVY := minY
		maxVY := (minY * -1) - 1

		var highest int
		for vx := 1; vx < maxX; vx++ {
			for step := 1; step <= maxX; step++ {
				x := X(vx, step)
				if x >= minX && x <= maxX {

					for vy := minVY; vy <= maxVY; vy++ {
						if passesWithin(vx, vy, minX, maxX, minY, maxY) {
							_, max := findMaxY(vx, vy)
							if max.y > highest {
								highest = max.y
							}
						}
					}
				}
			}
		}

		fmt.Printf("%d\n", highest)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func passesWithin(vx, vy, minX, maxX, minY, maxY int) bool {
	step := 0
	for {
		p := probePosition(step, vx, vy)
		if p.x > maxX || p.y < minY {
			return false
		} else if p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY {
			return true
		}
		step++
	}

	return false
}

func findMaxY(vx, vy int) (int, pos) {
	var step int
	max := pos{y: math.MinInt}
	for {
		p := probePosition(step, vx, vy)
		if p.y > max.y {
			max = p
		} else {
			break
		}

		step++
	}

	return step, max
}

func probePosition(step, vx, vy int) pos {
	return pos{
		x: X(vx, step),
		y: Y(vy, step),
	}
}

func X(vx, step int) int {
	var diff int
	if vx > 0 {
		diff = -1
	} else if vx < 0 {
		diff = 1
	}

	stepsBeforeZero := int(math.Min(float64(step), math.Abs(float64(vx))))

	return arithmeticSeries(vx, diff, stepsBeforeZero)
}

func Y(vy, step int) int {
	return arithmeticSeries(vy, -1, step)
}

func arithmeticSeries(start, diff, step int) int {
	return step * (2*start + (step-1)*diff) / 2
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
