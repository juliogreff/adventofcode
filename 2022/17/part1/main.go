package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juliogreff/adventofcode/pkg/mustparse"
	"github.com/juliogreff/adventofcode/pkg/mustread"
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
			tower   [][]bool
		)

		for scanner.Scan() {
			pattern = scanner.Text()
		}

		steps := 2022

		if len(os.Args) > 2 {
			steps = mustparse.Int(os.Args[2])
		}

		var nextRock, nextMove, total int

		const factor = 3

		x := 2
		y := len(tower) + factor
		rock := rocks[nextRock]

		for {
			var newX int
			switch pattern[nextMove] {
			case '>':
				newX = x + 1
			case '<':
				newX = x - 1
			}
			nextMove = (nextMove + 1) % len(pattern)

			if canMove(tower, rock, newX, y) {
				x = newX
			}

			cf := canFall(tower, rock, x, y)
			if cf {
				y--
			} else {
				tower = blit(tower, rock, x, y)

				total++

				nextRock = (nextRock + 1) % len(rocks)
				rock = rocks[nextRock]
				y = len(tower) + factor
				x = 2
			}

			if total == steps {
				break
			}
		}

		fmt.Printf("%d\n", len(tower))
	})
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

func blit(t [][]bool, rock rock, x, y int) [][]bool {
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
				}
			}
		}
	}

	return t
}

func printTower(t [][]bool) {
	for i := len(t) - 1; i >= 0; i-- {
		printRow(t[i])

		fmt.Println()
	}
}

func printRow(r []bool) {
	for _, b := range r {
		if b {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}
