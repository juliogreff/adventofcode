package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/juliogreff/adventofcode/pkg/intmath"
	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
)

const (
	factor   = 3
	maxRocks = 1000000000000
)

type rock [][]bool

var rocks = []rock{
	{{true, true, true, true}},
	{{false, true, false}, {true, true, true}, {false, true, false}},
	{{true, true, true}, {false, false, true}, {false, false, true}},
	{{true}, {true}, {true}, {true}},
	{{true, true}, {true, true}},
}

func main() {
	mustread.File(os.Args[1], func(scanner *bufio.Scanner) {
		var (
			pattern string
		)

		for scanner.Scan() {
			pattern = scanner.Text()
		}

		nRocks := maxRocks

		if len(os.Args) > 2 {
			nRocks = mustparse.Int(os.Args[2])
		}

		rocksUsed, height, seedRocks := simulate(pattern, nRocks)

		repetitions := (nRocks - seedRocks) / (rocksUsed - seedRocks)
		remainder := (nRocks - seedRocks) % (rocksUsed - seedRocks)

		_, seedHeight, _ := simulate(pattern, seedRocks)
		_, diffHeight, _ := simulate(pattern, seedRocks+remainder)

		lastRepeatHeight := heightAtRepeat(seedRocks, seedHeight, height, repetitions)

		fmt.Printf("%d\n", lastRepeatHeight+diffHeight-seedHeight)
	})
}

func heightAtRepeat(seedRocks, seedHeight, height, repetitions int) int {
	return seedHeight + repetitions*(height-seedHeight)
}

func simulate(pattern string, maxSteps int) (int, int, int) {
	var (
		seedHeight                int
		nextRock, nextMove, steps int
		min, max, accum           int
		tower                     [][]bool
		set                       [7]int
	)

	cache := make(map[string]int)

	x := 2
	y := len(tower) + factor
	rock := rocks[nextRock]

	for {
		k := key(set, nextMove, nextRock)

		if s, ok := cache[k]; ok {
			seedHeight = s
			break
		}

		var nx int
		switch pattern[nextMove] {
		case '>':
			nx = x + 1
		case '<':
			nx = x - 1
		}
		nextMove = (nextMove + 1) % len(pattern)

		if canMove(tower, rock, nx, y) {
			x = nx
		}

		cf := canFall(tower, rock, x, y)
		if cf {
			y--
		} else {
			tower, max = blit(tower, rock, x, y, max, &set)

			minSet := math.MaxInt
			for _, v := range set {
				if v < minSet {
					minSet = v
				}
			}

			if minSet > min {
				min = minSet
				tower = tower[min:]

				max -= min
				accum += min

				for k := range set {
					set[k] -= min
				}
			}

			cache[k] = steps

			steps++

			nextRock = (nextRock + 1) % len(rocks)
			rock = rocks[nextRock]
			y = len(tower) + factor
			x = 2
		}

		if steps == maxSteps {
			break
		}
	}

	return steps, accum + max, seedHeight
}

func canMove(t [][]bool, rock rock, x, y int) bool {
	if x < 0 || x+len(rock[0]) > 7 {
		return false
	}

	if y < len(t) {
		for i, row := range rock {
			if y+i >= len(t) {
				return true
			}

			for j, block := range row {
				if block && t[y+i][x+j] {
					return false
				}
			}
		}
	}

	return true
}

func canFall(t [][]bool, rock rock, x, y int) bool {
	if y == 0 {
		return false
	}

	for i, row := range rock {
		if y+i > len(t) {
			return true
		}

		for j, block := range row {
			if block && t[y-1+i][x+j] {
				return false
			}
		}
	}

	return true
}

func blit(t [][]bool, rock rock, x, y, max int, set *[7]int) ([][]bool, int) {
	for i, row := range rock {
		if y+i == len(t) {
			t = append(t, make([]bool, 7, 7))
		}

		for j, block := range row {
			if x+j < 0 || x+j > 6 {
				panic("out of bounds")
			} else if block {
				if t[y+i][x+j] {
					panic("overwriting")
				} else {
					t[y+i][x+j] = true

					val := y + i + 1
					set[x+j] = intmath.Max(val, set[x+j])
					max = intmath.Max(val, max)
				}
			}
		}
	}

	return t, max
}

func key(set [7]int, nextMove int, nextRock int) string {
	return fmt.Sprintf("%v-%d-%d", set, nextMove, nextRock)
}
